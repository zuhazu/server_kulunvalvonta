package data

import (
	"context"
	"goapi/internal/api/repository/models"
)

type DataService interface {
	Create(data *models.Student, ctx context.Context) error
	ReadOne(id int, ctx context.Context) (*models.Student, error)
	//ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Student, error)
	Update(data *models.Student, ctx context.Context) (int64, error)
	Delete(data *models.Student, ctx context.Context) (int64, error)
	ValidateData(data *models.Student) error
}

type DataError struct {
	Message string
}

func (de DataError) Error() string {
	return de.Message
}
