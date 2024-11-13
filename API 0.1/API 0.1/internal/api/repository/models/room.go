package models

import "context"

type Room struct {
	ID       int    `json:"id"`
	RoomID   int    `json:"room_id"`
	RoomName string `json:"room_name"`
}

type RoomRepository interface {
	Create(Data *Room, ctx context.Context) error
	ReadOne(id int, ctx context.Context) (*Room, error)
	// ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*Data, error)
	Update(data *Room, ctx context.Context) (int64, error)
	Delete(data *Room, ctx context.Context) (int64, error)
}
