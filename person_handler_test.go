package main

import (
	"context"
	"database/sql"
	"my-project/domain/repository"
	timeapihost "my-project/infrastructure/clients/timeapi"
	tulushost "my-project/infrastructure/clients/tulustech"

	"my-project/infrastructure/configuration"
	"my-project/infrastructure/logger"
	"my-project/infrastructure/persistence"
	"my-project/infrastructure/pubsub"
	"my-project/infrastructure/servicebus"
	"my-project/usecase"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	httpHandler "my-project/interfaces/http"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/sync/errgroup"
)

type PersonHandlerSuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository repository.IPerson
}

func (s *PersonHandlerSuite) SetupPersonHandlerSuite() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.repository = persistence.NewPersonRepository(db)
}

func (s *PersonHandlerSuite) TestGetCountrySuccess(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	_, ctx = errgroup.WithContext(ctx)

	db, _, _ := sqlmock.New()

	var (
		Name    = "Adam"
		Country = "Kuala Lumpur"
	)

	prep := s.mock.ExpectPrepare(regexp.QuoteMeta(`SELECT p.name, p.country 
	FROM persons AS p 
	WHERE u.name = ?`))
	prep.ExpectQuery().WithArgs("Adam").
		WillReturnRows(sqlmock.NewRows([]string{"name", "country"}).
			AddRow(Name, Country))

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

	userRepository := persistence.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	testUsecase := usecase.NewTestUsecase(tulusTechHost, testPubSub, testServiceBus, timeApiHost)

	personUsecase := usecase.NewPersonUsecase(s.repository)

	userHandler := httpHandler.NewUserHandler(userUsecase)
	personHandler := httpHandler.NewPersonHandler(personUsecase)
	testHandler := httpHandler.NewTestHandler(testUsecase)

	router := InitiateRouter(personHandler, userHandler, testHandler, userRepository)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/country/adam", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Kuala Lumpur", w.Body.String())
}

func (s *PersonHandlerSuite) TestGetCountryNotFound(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	_, ctx = errgroup.WithContext(ctx)

	db, _, _ := sqlmock.New()

	prep := s.mock.ExpectPrepare(regexp.QuoteMeta(`SELECT p.name, p.country 
	FROM persons AS p 
	WHERE u.name = ?`))
	prep.ExpectQuery().WithArgs("andi").WillReturnError(sql.ErrNoRows)

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

	userRepository := persistence.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	testUsecase := usecase.NewTestUsecase(tulusTechHost, testPubSub, testServiceBus, timeApiHost)

	personUsecase := usecase.NewPersonUsecase(s.repository)

	userHandler := httpHandler.NewUserHandler(userUsecase)
	personHandler := httpHandler.NewPersonHandler(personUsecase)
	testHandler := httpHandler.NewTestHandler(testUsecase)

	router := InitiateRouter(personHandler, userHandler, testHandler, userRepository)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/country/adam", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "Name not found", w.Body.String())
}
