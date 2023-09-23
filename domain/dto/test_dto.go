package dto

type TestDto struct {
	PubSub     string `json:"pub_sub"`
	ServiceBus string `json:"service_bus"`
	TulusTech  string `json:"tulus_tech"`
}

type ReqUriParamTimeApi struct {
	Timezone string `json:"timezone" binding:"required" uri:"timezone"`
}
