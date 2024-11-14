package data

import (
	"context"
	"goapi/internal/api/repository/models"
)

type PersonService interface {
	CreatePerson(data *models.Person, ctx context.Context) error
	ReadOnePerson(id int, ctx context.Context) (*models.Person, error)
	ValidateData(data *models.Person) error
}

type PersonError struct {
	Message string
}

func (de PersonError) Error() string {
	return de.Message
}
