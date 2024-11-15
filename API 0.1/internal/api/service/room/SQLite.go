package data

import (
	"context"
	"goapi/internal/api/repository/models"
)

type RoomServiceSQLite struct {
	repo models.RoomRepository
}

func NewRoomServiceSQLite(repo models.RoomRepository) *RoomServiceSQLite {
	return &RoomServiceSQLite{
		repo: repo,
	}
}

func (ds *RoomServiceSQLite) CreateRoom(data *models.Room, ctx context.Context) error {

	if err := ds.ValidateData(data); err != nil {
		return RoomError{Message: "Invalid room data."}
	}
	return ds.repo.CreateRoom(data, ctx)
}

func (ds *RoomServiceSQLite) ReadOneRoom(id int, ctx context.Context) (*models.Room, error) {

	data, err := ds.repo.ReadOneRoom(id, ctx)
	if err != nil {
		return nil, err
	}

	_ = data

	return data, nil
}

func (ds *RoomServiceSQLite) ValidateData(data *models.Room) error {
	var errMsg string
	if errMsg != "" {
		return RoomError{Message: errMsg}
	}
	return nil
}
