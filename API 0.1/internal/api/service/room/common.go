package data

import (
	"context"
	"goapi/internal/api/repository/models"
)

type RoomService interface {
	//Luodaan room
	CreateRoom(data *models.Room, ctx context.Context) error
	//Haetaan huone
	ReadOneRoom(id int, ctx context.Context) (*models.Room, error)
	//Validoidaan data
	ValidateData(data *models.Room) error
	//Haetaan person-entiteetit roomID:n perusteella
	GetPersonsByRoomID(room_id string, ctx context.Context) ([]*models.Person, error)
	//Tarkistetaan onko huone olemassa
	CheckIfRoomExist(room_id string, ctx context.Context) (bool, error)
}

type RoomError struct {
	Message string
}

func (re RoomError) Error() string {
	return re.Message
}
