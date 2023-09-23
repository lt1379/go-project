package usecase

import (
	"context"
	"my-project/domain/model"
	"my-project/domain/repository"
	"my-project/infrastructure/logger"
)

type IPersonUsecase interface {
	GetByPersonName(ctx context.Context, req model.ReqPerson) (string, error)
}

type PersonUsecase struct {
	personRepository repository.IPerson
}

func NewPersonUsecase(personRepository repository.IPerson) IPersonUsecase {
	return &PersonUsecase{personRepository: personRepository}
}

func (personUsecase *PersonUsecase) GetByPersonName(ctx context.Context, req model.ReqPerson) (string, error) {
	var res string

	person, err := personUsecase.personRepository.GetByName(ctx, req.Name)
	if err != nil {
		logger.GetLogger().WithField("error", err).Error("Error while Getting personname")
		return res, err
	}

	res = person.Country

	return res, nil
}
