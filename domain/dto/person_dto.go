package dto

type ReqUriParamPerson struct {
	Name string `json:"name" binding:"required"`
}
