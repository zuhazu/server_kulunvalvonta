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
	// Luodaan uusi henkilö
	CreatePerson(Data *Person, ctx context.Context) error
	// Luetaan yksi henkilö
	ReadOnePerson(id int, ctx context.Context) (*Person, error)
	// Päivitetään yhtä henkilöä
	UpdatePerson(data *Person, ctx context.Context) (int64, error)
	// Poistetaan yksi henkilö
	DeletePerson(data *Person, ctx context.Context) (int64, error)
	// Päivitetään henkilön roomId henkilön tag id:n ja arduinon luokkaid:n perusteella
	UpdateRoomIDByTagID(tagID string, newRoomID string, ctx context.Context) (string, error)
	// Haetaan henkilöt roomId:n perusteella
	ReadPersonsByRoomId(roomId string, ctx context.Context) ([]*Person, error)
}
