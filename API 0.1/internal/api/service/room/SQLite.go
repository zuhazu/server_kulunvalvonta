package room

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

// Luodaan room
func (ds *RoomServiceSQLite) CreateRoom(data *models.Room, ctx context.Context) error {
	exist, err := ds.CheckIfRoomExist(data.RoomID, ctx)
	if err != nil {
		return err
	}
	if exist {
		return RoomError{"Room does not exist"}
	}

	if err := ds.ValidateData(data); err != nil {
		return RoomError{"Invalid data"}
	}
	return ds.repo.CreateRoom(data, ctx)
}

// Haetaan huone
func (ds *RoomServiceSQLite) ReadOneRoom(id int, ctx context.Context) (*models.Room, error) {

	data, err := ds.repo.ReadOneRoom(id, ctx)
	if err != nil {
		return nil, err
	}

	_ = data

	return data, nil
}

// Validoidaan data
func (ds *RoomServiceSQLite) ValidateData(data *models.Room) error {
	var errMsg string
	if errMsg != "" {
		return RoomError{Message: errMsg}
	}
	return nil
}

// Haetaan person-entiteetit huoneID:n perusteella
func (rs *RoomServiceSQLite) GetPersonsByRoomID(room_id string, ctx context.Context) ([]*models.Person, error) {
	exist, err := rs.CheckIfRoomExist(room_id, ctx)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, RoomError{"Invalid room"}
	}

	data, err := rs.repo.GetPersonsByRoomID(room_id, ctx)
	if err != nil {
		return nil, err
	}

	_ = data

	return data, nil
}

// Tarkistetaan onko huone olemassa
func (rs *RoomServiceSQLite) CheckIfRoomExist(room_id string, ctx context.Context) (bool, error) {
	data, err := rs.repo.ReadOneRoomByRoomID(room_id, ctx)
	if err != nil {
		return false, err
	}
	if data != nil {
		return true, nil
	}
	return false, nil
}
