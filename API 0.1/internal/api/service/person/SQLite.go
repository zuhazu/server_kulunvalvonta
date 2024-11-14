package data

import (
	"context"
	"goapi/internal/api/repository/models"
)

type PersonServiceSQLite struct {
	repo models.PersonRepository
}

func NewPersonServiceSQLite(repo models.PersonRepository) *PersonServiceSQLite {
	return &PersonServiceSQLite{
		repo: repo,
	}
}

func (ds *PersonServiceSQLite) CreatePerson(data *models.Person, ctx context.Context) error {

	if err := ds.ValidateData(data); err != nil {
		return PersonError{Message: "InvalMockDataServiceSuccessfulid data."}
	}
	return ds.repo.CreatePerson(data, ctx)
}

func (ds *PersonServiceSQLite) ReadOnePerson(id int, ctx context.Context) (*models.Person, error) {

	data, err := ds.repo.ReadOnePerson(id, ctx)
	if err != nil {
		return nil, err
	}

	_ = data

	return data, nil
}

func (ds *PersonServiceSQLite) ValidateData(data *models.Person) error {
	var errMsg string
	if errMsg != "" {
		return PersonError{Message: errMsg}
	}
	return nil
}
