package models

import "context"

type Room struct {
	ID       int    `json:"id"`
	RoomID   string `json:"room_id"`
	RoomName string `json:"room_name"`
}

type RoomRepository interface {
	CreateRoom(Data *Room, ctx context.Context) error
	ReadOneRoom(id int, ctx context.Context) (*Room, error)
	GetPersonsByRoomID(room_id string, ctx context.Context) ([]*Person, error)
	ReadOneRoomByRoomID(room_id string, ctx context.Context) (*Room, error)
}
