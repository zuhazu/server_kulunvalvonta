package models

import "context"

type Room struct {
	ID       int    `json:"id"`
	RoomID   string `json:"room_id"`
	RoomName string `json:"room_name"`
}

type RoomRepository interface {
	//Luodaan room
	CreateRoom(Data *Room, ctx context.Context) error
	//Haetaan yksi huone
	ReadOneRoom(id int, ctx context.Context) (*Room, error)
	//Haetaan personit roomId:n perusteella
	GetPersonsByRoomID(room_id string, ctx context.Context) ([]*Person, error)
	//haetaan room roomID:n perusteella
	ReadOneRoomByRoomID(room_id string, ctx context.Context) (*Room, error)
}
