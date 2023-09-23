package usecase

import (
	"context"
	"encoding/json"
	"my-project/domain/dto"
	"my-project/infrastructure/clients/timeapi"
	timeapimodels "my-project/infrastructure/clients/timeapi/models"
	tulushost "my-project/infrastructure/clients/tulustech"
	"my-project/infrastructure/clients/tulustech/models"
	"my-project/infrastructure/logger"
	"my-project/infrastructure/pubsub"
	"my-project/infrastructure/servicebus"
)

type ITestUsecase interface {
	Test(ctx context.Context) dto.TestDto
	GetCurrentTime(ctx context.Context, timeZone string) (dto.ResTimeApiDto, error)
}

type TestUsecase struct {
	TulusTechHost  tulushost.ITulusHost
	TestPubSub     pubsub.ITestPubSub
	TestServiceBus servicebus.ITestServiceBus
	TimeApiHost    timeapi.ITimeApiHost
}

func NewTestUsecase(tulusTechHost tulushost.ITulusHost, testPubSub pubsub.ITestPubSub, testServiceBus servicebus.ITestServiceBus, timeApiHost timeapi.ITimeApiHost) ITestUsecase {
	return &TestUsecase{TulusTechHost: tulusTechHost, TestPubSub: testPubSub, TestServiceBus: testServiceBus, TimeApiHost: timeApiHost}
}

func (testUsecase *TestUsecase) Test(ctx context.Context) dto.TestDto {
	var res dto.TestDto

	res.PubSub = "Not OK"
	res.ServiceBus = "Not OK"

	msg := "Hello"
	byteMsg, err := json.Marshal(msg)
	if err != nil {
		logger.GetLogger().Error("Error while marshalling")
		return res
	}
	publishResponse, err := testUsecase.TestPubSub.Publish(ctx, "topic", byteMsg)
	if err != nil {
		logger.GetLogger().Error("Error while publishing message")
		res.PubSub = err.Error()
		return res
	}
	logger.GetLogger().WithField("publishResponse", publishResponse).Info("Successfully published")
	res.PubSub = "OK"

	err = testUsecase.TestServiceBus.SendMessage(byteMsg)
	if err != nil {
		logger.GetLogger().Error("Error while publishing message with service bus")
		res.ServiceBus = err.Error()
		return res
	}
	res.ServiceBus = "OK"

	reqHeader := models.ReqHeader{}
	randomTypingRes, err := testUsecase.TulusTechHost.GetRandomTyping(ctx, reqHeader)
	if err != nil {
		logger.GetLogger().Error("Error while get random typing")
		res.ServiceBus = err.Error()
		return res
	}
	logger.GetLogger().WithField("randomTypingResponse", randomTypingRes).Info("Successfully get random typing")
	res.TulusTech = "OK"

	return res
}

func (testUsecase *TestUsecase) GetCurrentTime(ctx context.Context, timeZone string) (dto.ResTimeApiDto, error) {
	var res dto.ResTimeApiDto

	reqHeader := timeapimodels.ReqHeader{}
	resTimeApi, err := testUsecase.TimeApiHost.GetCurrentTime(ctx, reqHeader, timeZone)
	if err != nil {
		logger.GetLogger().Error("Error while get timezone")
		return res, err
	}
	logger.GetLogger().WithField("resTimeApiponse", resTimeApi).Info("Successfully get timezone")
	resTimeApiByte, err := json.Marshal(resTimeApi)
	if err != nil {
		logger.GetLogger().Error("Error while marshalling")
		return res, err
	}

	if err := json.Unmarshal(resTimeApiByte, &res); err != nil {
		logger.GetLogger().Error("Error while unmarshalling")
		return res, err
	}

	return res, nil
}
