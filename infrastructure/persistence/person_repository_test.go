package persistence

import (
	"context"
	"database/sql"
	"my-project/domain/model"
	"my-project/domain/repository"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type PersonSuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository repository.IPerson
}

func (s *PersonSuite) SetupPersonSuite() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.repository = NewPersonRepository(db)
}

func TestPersonRepository(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *PersonSuite) TestPersonRepository_GetByName() {
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

	res, err := s.repository.GetByName(context.Background(), "Adam")
	exp := model.Person{
		Name:    Name,
		Country: Country,
	}

	require.Nil(s.T(), err)
	require.Equal(s.T(), exp, res)
}
