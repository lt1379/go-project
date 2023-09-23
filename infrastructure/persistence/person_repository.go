package persistence

import (
	"context"
	"database/sql"

	"my-project/domain/model"
	"my-project/domain/repository"
	"my-project/infrastructure/logger"
)

type PersonRepository struct {
	sqlDB *sql.DB
}

func NewPersonRepository(sqlDB *sql.DB) repository.IPerson {
	return &PersonRepository{sqlDB}
}

func (personRepository *PersonRepository) GetByName(ctx context.Context, name string) (model.Person, error) {
	var person model.Person

	statement, err := personRepository.sqlDB.PrepareContext(ctx, `SELECT p.name, p.country 
	FROM persons AS p 
	WHERE p.name = ?`)

	if err != nil {
		logger.GetLogger().WithField("error", err).Error("Error while prepare statement")
		return person, err
	}
	defer statement.Close()

	result := statement.QueryRow(name)
	err = result.Scan(&person.Country)
	if err != nil {
		logger.GetLogger().WithField("error", err).Error("Error while query")
		return person, err
	}

	return person, nil
}
