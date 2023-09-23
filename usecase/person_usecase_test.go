package usecase_test

import (
	"context"
	"database/sql"
	"my-project/domain/model"
	"my-project/mocks/repomocks"
	"my-project/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersonUsecase_GetByNameSuccess(t *testing.T) {
	personRepository := &repomocks.IPerson{}
	personRepository.On("GetByName", context.Background(), "Adam").Return(model.Person{
		Name:    "Adam",
		Country: "Kuala Lumpur",
	}, nil).Once()

	personUsecase := usecase.NewPersonUsecase(personRepository)
	country, err := personUsecase.GetByPersonName(context.Background(), model.ReqPerson{
		Name: "Adam",
	})

	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.Equal(t, "Kuala Lumpur", country)
}

func TestPersonUsecase_GetByNameError(t *testing.T) {
	personRepository := &repomocks.IPerson{}
	personRepository.On("GetByName", context.Background(), "Adam").Return(model.Person{}, sql.ErrNoRows).Once()

	personUsecase := usecase.NewPersonUsecase(personRepository)
	country, err := personUsecase.GetByPersonName(context.Background(), model.ReqPerson{
		Name: "Adam",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "", country)
}
