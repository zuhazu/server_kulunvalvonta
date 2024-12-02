package person

import (
	"context"
	"goapi/internal/api/repository/models"
)

type PersonService interface {
	// Luodaan person
	CreatePerson(data *models.Person, ctx context.Context) error
	// Haetaan person
	ReadOnePerson(id int, ctx context.Context) (*models.Person, error)
	// Päivitetään person
	UpdatePerson(data *models.Person, ctx context.Context) (int64, error)
	// Poistetaan person
	DeletePerson(data *models.Person, ctx context.Context) (int64, error)
	// Päivitetään roomID tagID:n perusteella
	UpdateRoomIDByTagID(tagID, newRoomID string, ctx context.Context) (string, error)
	// Valitoidaan data
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
