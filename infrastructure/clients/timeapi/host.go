package timeapi

import (
	"context"
	"encoding/json"
	"fmt"
	"my-project/infrastructure/clients"
	"my-project/infrastructure/clients/timeapi/models"
)

type ITimeApiHost interface {
	GetCurrentTime(ctx context.Context, reqHeader models.ReqHeader, timeZone string) (models.ResTimeApi, error)
}

type TimeApiHost struct {
	id   string
	host string
}

func NewTimeApiHost(host string) ITimeApiHost {
	return &TimeApiHost{host: host}
}

func (TimeApiHost *TimeApiHost) GetCurrentTime(ctx context.Context, reqHeader models.ReqHeader, timeZone string) (models.ResTimeApi, error) {
	var res models.ResTimeApi

	endpoint := fmt.Sprintf("/api/Time/current/zone?timeZone=%s", timeZone)
	method := "GET"

	reqMapHeader := map[string]string{
		"Accept":       reqHeader.Accept,
		"Content-Type": reqHeader.ContentType,
		"Cookie":       reqHeader.Cookie,
	}
	hostClient := clients.NewHost(TimeApiHost.host, endpoint, method, nil, reqMapHeader, nil)
	byteData, statusCode, err := hostClient.HTTPGet()
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(byteData, &res); err != nil {
		return res, err
	}

	if statusCode < 200 && statusCode > 299 {
		return res, fmt.Errorf("something occurred with server")
	}

	return res, nil
}
