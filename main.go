package main

import (
	"context"
	"fmt"
	"log"
	timeapihost "my-project/infrastructure/clients/timeapi"
	tulushost "my-project/infrastructure/clients/tulustech"
	"my-project/infrastructure/configuration"
	"my-project/infrastructure/logger"
	"my-project/infrastructure/persistence"
	"my-project/infrastructure/pubsub"
	"my-project/infrastructure/servicebus"
	httpHandler "my-project/interfaces/http"
	"my-project/usecase"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	httpServer *http.Server
)

func recoverPanic() {
	if err := recover(); err != nil {
		fmt.Printf("RECOVERED: %v\n", err)
	}
}

func main() {
	defer recoverPanic()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	g, ctx := errgroup.WithContext(ctx)

	// configuration.LoadConfig()

	app := configuration.C.App

	mysqlDb, psqlDb, err := InitiateDatabase()
	if err != nil {
		panic(err)
	}

	logger.GetLogger().WithField("MySQLDb", mysqlDb.Ping()).WithField("PSQLDb", psqlDb.Ping()).Info("Database connected.")

	pubSubClient, err := pubsub.NewPubSub(ctx, configuration.C.Pubsub.ProjectID)
	if err != nil {
		logger.GetLogger().WithField("error", err).Error("Error while instantiate PubSub")
		panic(err)
	}

	azServiceBusClient, err := servicebus.NewServiceBus(ctx, configuration.C.ServiceBus.Namespace)
	if err != nil {
		logger.GetLogger().WithField("error", err).Error("Error while instantiate ServiceBus")
		panic(err)
	}

	tulusTechHost := tulushost.NewTulusHost(configuration.C.TulusTech.Host)
	timeApiHost := timeapihost.NewTimeApiHost(configuration.C.TimeApi.Host)

	testPubSub := pubsub.NewTestPubSub(pubSubClient)
	testServiceBus := servicebus.NewTestServiceBus(azServiceBusClient)

	userRepository := persistence.NewUserRepository(psqlDb)
	userUsecase := usecase.NewUserUsecase(userRepository)
	testUsecase := usecase.NewTestUsecase(tulusTechHost, testPubSub, testServiceBus, timeApiHost)

	personRepository := persistence.NewPersonRepository(mysqlDb)
	personUsecase := usecase.NewPersonUsecase(personRepository)

	userHandler := httpHandler.NewUserHandler(userUsecase)
	personHandler := httpHandler.NewPersonHandler(personUsecase)
	testHandler := httpHandler.NewTestHandler(testUsecase)

	router := InitiateRouter(personHandler, userHandler, testHandler, userRepository)

	if err != nil {
		logger.GetLogger().WithField("error", err).Error("Error while StartSubscription")
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	port := app.Port
	logger.GetLogger().WithField("port", port).Info("Starting application")
	g.Go(func() error {
		httpServer := &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			Handler:      router,
			ReadTimeout:  0,
			WriteTimeout: 0,
		}
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		logger.GetLogger().WithField("port", port).Error("Application start")
		return nil
	})

	select {
	case <-interrupt:
		fmt.Println("Exit")
		os.Exit(1)
		break
	case <-ctx.Done():
		break
	}

	cancel()
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if httpServer != nil {
		_ = httpServer.Shutdown(shutdownCtx)
	}

	err = g.Wait()
	if err != nil {
		log.Printf("server returning an error %v", err)
		os.Exit(2)
	}
}
