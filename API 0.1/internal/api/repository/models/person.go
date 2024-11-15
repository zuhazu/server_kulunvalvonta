package models

import "context"

type Person struct {
	ID         int    `json:"id"`
	PersonID   string `json:"person_id"`
	TagID      string `json:"tag_id"`
	PersonName string `json:"person_name"`
	RoomID     string `json:"room_id"`
}

type PersonRepository interface {
	CreatePerson(Data *Person, ctx context.Context) error
	ReadOnePerson(id int, ctx context.Context) (*Person, error)
	//ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*Data, error)
	UpdatePerson(data *Person, ctx context.Context) (int64, error)
	DeletePerson(data *Person, ctx context.Context) (int64, error)
	UpdateRoomIDByTagID(personID string, tagID string, newRoomID string, ctx context.Context) (string, error)
}
