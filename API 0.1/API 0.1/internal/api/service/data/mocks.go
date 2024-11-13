package data

import (
	"context"
	"goapi/internal/api/repository/models"
)

// * Mock implementation of DataService for testing purposes, always returns a successful response and Data object(s) *
type MockDataServiceSuccessful struct{}

func (m *MockDataServiceSuccessful) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Student, error) {
	return []*models.Student{
		{
			ID:          1,
			StudentID:   121,
			StudentName: "device1",
			RoomID:      6,
		},
		{
			ID:          2,
			StudentID:   121,
			StudentName: "device1",
			RoomID:      6,
		},
	}, nil
}

func (m *MockDataServiceSuccessful) ReadOne(id int, ctx context.Context) (*models.Student, error) {
	return &models.Student{
		ID:          1,
		StudentID:   121,
		StudentName: "device1",
		RoomID:      6,
	}, nil
}

func (m *MockDataServiceSuccessful) Create(data *models.Student, ctx context.Context) error {
	return nil
}

func (m *MockDataServiceSuccessful) Update(data *models.Student, ctx context.Context) (int64, error) {
	return 1, nil
}

func (m *MockDataServiceSuccessful) Delete(data *models.Student, ctx context.Context) (int64, error) {
	return 1, nil
}

func (m *MockDataServiceSuccessful) ValidateData(data *models.Student) error {
	return nil
}

// * Mock implementation of DataService for testing purposes, always returns empty data *

type MockDataServiceNotFound struct{}

func (m *MockDataServiceNotFound) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Student, error) {
	return []*models.Student{}, nil
}

func (m *MockDataServiceNotFound) ReadOne(id int, ctx context.Context) (*models.Student, error) {
	return nil, nil
}

func (m *MockDataServiceNotFound) Create(data *models.Student, ctx context.Context) error {
	return nil
}

func (m *MockDataServiceNotFound) Update(data *models.Student, ctx context.Context) (int64, error) {
	return 0, nil
}

func (m *MockDataServiceNotFound) Delete(data *models.Student, ctx context.Context) (int64, error) {
	return 0, nil
}

func (m *MockDataServiceNotFound) ValidateData(data *models.Student) error {
	return nil
}

// * Mock implementation of DataService for testing purposes, always returns an error *
type MockDataServiceError struct{}

func (m *MockDataServiceError) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Student, error) {
	return nil, DataError{Message: "Error reading data."}
}

func (m *MockDataServiceError) ReadOne(id int, ctx context.Context) (*models.Student, error) {
	return nil, DataError{Message: "Error reading data."}
}

func (m *MockDataServiceError) Create(data *models.Student, ctx context.Context) error {
	return DataError{Message: "Error creating data."}
}

func (m *MockDataServiceError) Update(data *models.Student, ctx context.Context) (int64, error) {
	return 0, DataError{Message: "Error updating data."}
}

func (m *MockDataServiceError) Delete(data *models.Student, ctx context.Context) (int64, error) {
	return 0, DataError{Message: "Error deleting data."}
}

func (m *MockDataServiceError) ValidateData(data *models.Student) error {
	return nil
}
