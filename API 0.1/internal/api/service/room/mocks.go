package room

import (
	"context"
	"goapi/internal/api/repository/models"
)

// * Mock implementation of DataService for testing purposes, always returns a successful response and Data object(s) *
type MockRoomServiceSuccessful struct{}

func (m *MockRoomServiceSuccessful) CreateRoom(data *models.Room, ctx context.Context) error {
	return nil
}

func (m *MockRoomServiceSuccessful) ReadOneRoom(id int, ctx context.Context) (*models.Room, error) {
	return &models.Room{
		ID:       2,
		RoomID:   "12",
		RoomName: "kakkuli",
	}, nil
}

func (m *MockRoomServiceSuccessful) ValidateData(data *models.Room) error {
	return nil
}

func (m *MockRoomServiceSuccessful) GetPersonsByRoomID(roomId string, ctx context.Context) ([]*models.Person, error) {
	return []*models.Person{}, nil
	// return []*models.Person{
	// 	{ID: 1, PersonID: "P001", TagID: "T001", PersonName: "Alice McDonals", RoomID: "100100"},
	// 	{ID: 2, PersonID: "P002", TagID: "T002", PersonName: "Bob Hesburger", RoomID: "100100"},
	// }, nil
}

func (m *MockRoomServiceSuccessful) CheckIfRoomExist(room_id string, ctx context.Context) (bool, error) {
	return false, nil
}

// * Mock implementation of DataService for testing purposes, always returns empty data *

type MockRoomServiceNotFound struct{}

func (m *MockRoomServiceNotFound) CreateRoom(data *models.Room, ctx context.Context) error {
	return nil
}

func (m *MockRoomServiceNotFound) ReadOneRoom(id int, ctx context.Context) (*models.Room, error) {
	return nil, nil
}

func (m *MockRoomServiceNotFound) ValidateData(data *models.Room) error {
	return nil
}

func (m *MockRoomServiceNotFound) GetPersonsByRoomID(roomId string, ctx context.Context) ([]*models.Person, error) {
	return nil, RoomError{"Page not found"}
}

func (m *MockRoomServiceNotFound) CheckIfRoomExist(room_id string, ctx context.Context) (bool, error) {
	return false, nil
}

// * Mock implementation of DataService for testing purposes, always returns an error *
type MockRoomServiceError struct{}

func (m *MockRoomServiceError) CreateRoom(data *models.Room, ctx context.Context) error {
	return nil
}

func (m *MockRoomServiceError) ReadOneRoom(id int, ctx context.Context) (*models.Room, error) {
	return nil, RoomError{Message: "Error reading data."}
}

func (m *MockRoomServiceError) ValidateData(data *models.Room) error {
	return nil
}

func (m *MockRoomServiceError) GetPersonsByRoomID(roomId string, ctx context.Context) ([]*models.Person, error) {
	return []*models.Person{}, nil
}

func (m *MockRoomServiceError) CheckIfRoomExist(room_id string, ctx context.Context) (bool, error) {
	return false, nil
}
