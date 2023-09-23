package http

import (
	"my-project/infrastructure/logger"
	"my-project/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ITestHandler interface {
	Test(c *gin.Context)
	GetCurrentTime(c *gin.Context)
}

type TestHandler struct {
	TestUsecase usecase.ITestUsecase
}

func NewTestHandler(testUsecase usecase.ITestUsecase) ITestHandler {
	return &TestHandler{TestUsecase: testUsecase}
}

func (testHandler *TestHandler) Test(c *gin.Context) {
	res := testHandler.TestUsecase.Test(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func (testHandler *TestHandler) GetCurrentTime(c *gin.Context) {
	// var req dto.ReqUriParamTimeApi

	// if err := c.ShouldBindUri(&req); err != nil {
	// 	logger.GetLogger().WithField("error", err).Error("An error occurred")
	// 	c.JSON(http.StatusBadRequest, "Please provide a valid timezone")
	// 	return
	// }

	timezone := c.Query("timezone")
	timezone = cases.Title(language.Indonesian, cases.NoLower).String(timezone)
	res, err := testHandler.TestUsecase.GetCurrentTime(c.Request.Context(), timezone)
	if err != nil {
		logger.GetLogger().WithField("error", err).Error("An error occurred")
		c.JSON(http.StatusBadRequest, "Timezone not found")
		return
	}

	c.JSON(http.StatusOK, res)
}
