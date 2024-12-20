package data

import (
	"context"
	"goapi/internal/api/repository/models"
)

type RoomService interface {
	CreateRoom(data *models.Room, ctx context.Context) error
	ReadOneRoom(id int, ctx context.Context) (*models.Room, error)
	ValidateData(data *models.Room) error
}

type RoomError struct {
	Message string
}

func (re RoomError) Error() string {
	return re.Message
}
