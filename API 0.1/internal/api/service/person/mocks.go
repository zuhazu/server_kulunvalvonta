package person

import (
	"context"
	"goapi/internal/api/repository/models"
)

// * Mock implementation of DataService for testing purposes, always returns a successful response and Data object(s) *
type MockPersonServiceSuccessful struct{}

// func (m *MockPersonServiceSuccessful) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Person, error) {
// 	return []*models.Person{
// 		{
// 			ID:         1,
// 			PersonID:   "121",
// 			TagID:      "221",
// 			PersonName: "Pekka Puupää",
// 			RoomID:     "036",
// 		},
// 		{
// 			ID:         2,
// 			PersonID:   "212",
// 			TagID:      "222",
// 			PersonName: "Matti Muumaa",
// 			RoomID:     "035",
// 		},
// 	}, nil
// }

func (m *MockPersonServiceSuccessful) ReadOnePerson(id int, ctx context.Context) (*models.Person, error) {
	return &models.Person{
		ID:         1,
		PersonID:   "121",
		TagID:      "221",
		PersonName: "Pekka Puupää",
		RoomID:     "036",
	}, nil
}

func (m *MockPersonServiceSuccessful) CreatePerson(data *models.Person, ctx context.Context) error {
	return nil
}

func (m *MockPersonServiceSuccessful) UpdatePerson(data *models.Person, ctx context.Context) (int64, error) {
	return 1, nil
}

func (m *MockPersonServiceSuccessful) DeletePerson(data *models.Person, ctx context.Context) (int64, error) {
	return 1, nil
}

func (m *MockPersonServiceSuccessful) UpdateRoomIDByTagID(tagID string, newRoomID string, ctx context.Context) (string, error) {
	return "success", nil
}

func (m *MockPersonServiceSuccessful) ReadPersonsByRoomId(roomId string, ctx context.Context) ([]*models.Person, error) {
	people := []*models.Person{
		{ID: 1, PersonID: "P001", TagID: "T001", PersonName: "Alice McDonals", RoomID: "100100"},
		{ID: 2, PersonID: "P002", TagID: "T002", PersonName: "Bob Hesburger", RoomID: "100100"},
	}
	return people, nil
}

func (m *MockPersonServiceSuccessful) ValidateData(data *models.Person) error {
	return nil
}

// * Mock implementation of DataService for testing purposes, always returns empty data *

type MockPersonServiceNotFound struct{}

// func (m *MockPersonServiceNotFound) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Person, error) {
// 	return []*models.Person{}, nil
// }

func (m *MockPersonServiceNotFound) ReadOnePerson(id int, ctx context.Context) (*models.Person, error) {
	return nil, nil
}

func (m *MockPersonServiceNotFound) CreatePerson(data *models.Person, ctx context.Context) error {
	return nil
}

func (m *MockPersonServiceNotFound) UpdatePerson(data *models.Person, ctx context.Context) (int64, error) {
	return 0, nil
}

func (m *MockPersonServiceNotFound) DeletePerson(data *models.Person, ctx context.Context) (int64, error) {
	return 0, nil
}

func (m *MockPersonServiceNotFound) UpdateRoomIDByTagID(tagID string, newRoomID string, ctx context.Context) (string, error) {
	return "success", nil
}

func (m *MockPersonServiceNotFound) ReadPersonsByRoomId(roomId string, ctx context.Context) ([]*models.Person, error) {
	people := []*models.Person{
		{ID: 1, PersonID: "P001", TagID: "T001", PersonName: "Alice McDonals", RoomID: "100100"},
		{ID: 2, PersonID: "P002", TagID: "T002", PersonName: "Bob Hesburger", RoomID: "100100"},
	}
	return people, nil
}

func (m *MockPersonServiceNotFound) ValidateData(data *models.Person) error {
	return nil
}

// * Mock implementation of DataService for testing purposes, always returns an error *
type MockPersonServiceError struct{}

// func (m *MockPersonServiceError) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Person, error) {
// 	return nil, PersonError{Message: "Error reading data."}
// }

func (m *MockPersonServiceError) ReadOnePerson(id int, ctx context.Context) (*models.Person, error) {
	return nil, PersonError{Message: "Error reading data."}
}

func (m *MockPersonServiceError) CreatePerson(data *models.Person, ctx context.Context) error {
	return PersonError{Message: "Error creating data."}
}

func (m *MockPersonServiceError) UpdatePerson(data *models.Person, ctx context.Context) (int64, error) {
	return 0, PersonError{Message: "Error updating data."}
}

func (m *MockPersonServiceError) DeletePerson(data *models.Person, ctx context.Context) (int64, error) {
	return 0, PersonError{Message: "Error deleting data."}
}

func (m *MockPersonServiceError) UpdateRoomIDByTagID(tagID string, newRoomID string, ctx context.Context) (string, error) {
	return "success", nil
}

func (m *MockPersonServiceError) ReadPersonsByRoomId(roomId string, ctx context.Context) ([]*models.Person, error) {
	people := []*models.Person{
		{ID: 1, PersonID: "P001", TagID: "T001", PersonName: "Alice McDonals", RoomID: "100100"},
		{ID: 2, PersonID: "P002", TagID: "T002", PersonName: "Bob Hesburger", RoomID: "100100"},
	}
	return people, nil
}

func (m *MockPersonServiceError) ValidateData(data *models.Person) error {
	return nil
}
