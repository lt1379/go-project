package http

import (
	"fmt"
	"my-project/domain/dto"
	"my-project/domain/model"
	"my-project/infrastructure/logger"
	"my-project/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type IPersonHandler interface {
	GetCountry(c *gin.Context)
}

type PersonHandler struct {
	personUsecase usecase.IPersonUsecase
}

func NewPersonHandler(personUsecase usecase.IPersonUsecase) IPersonHandler {
	return &PersonHandler{personUsecase: personUsecase}
}

func (personHandler *PersonHandler) GetCountry(c *gin.Context) {
	var req dto.ReqUriParamPerson

	if err := c.ShouldBindUri(&req); err != nil {
		logger.GetLogger().WithField("error", err).Error("An error occurred")
		c.JSON(http.StatusBadRequest, "Please provide a valid name")
		return
	}

	reqPerson := model.ReqPerson{}
	reqPerson.Name = cases.Title(language.Indonesian, cases.NoLower).String(req.Name)
	fmt.Println(reqPerson.Name)
	res, err := personHandler.personUsecase.GetByPersonName(c.Request.Context(), reqPerson)
	if err != nil {
		logger.GetLogger().WithField("error", err).Error("An error occurred")
		c.JSON(http.StatusBadRequest, "Name not found")
		return
	}

	c.JSON(http.StatusOK, res)
}
