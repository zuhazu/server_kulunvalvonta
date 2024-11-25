package person

import (
	"context"
	"goapi/internal/api/repository/models"
)

type PersonService interface {
	CreatePerson(data *models.Person, ctx context.Context) error
	ReadOnePerson(id int, ctx context.Context) (*models.Person, error)
	UpdatePerson(data *models.Person, ctx context.Context) (int64, error)
	DeletePerson(data *models.Person, ctx context.Context) (int64, error)
	UpdateRoomIDByTagID(tagID, newRoomID string, ctx context.Context) (string, error)
	ValidateData(data *models.Person) error
	// Haetaan henkilöt roomId:n perusteella
	ReadPersonsByRoomId(roomId string, ctx context.Context) ([]*models.Person, error)
}

type PersonError struct {
	Message string
}

func (de PersonError) Error() string {
	return de.Message
}
