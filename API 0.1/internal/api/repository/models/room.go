package models

import "context"

type Room struct {
	ID       int    `json:"id"`
	RoomID   int    `json:"room_id"`
	RoomName string `json:"room_name"`
}

type RoomRepository interface {
	CreateRoom(Data *Room, ctx context.Context) error
	ReadOneRoom(id int, ctx context.Context) (*Room, error)
}
