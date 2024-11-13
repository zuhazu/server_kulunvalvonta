package data

import (
	"context"
	"goapi/internal/api/repository/models"
)

// * Implementation of DataService for SQLite database *
type DataServiceSQLite struct {
	repo models.StudentRepository
}

func NewDataServiceSQLite(repo models.StudentRepository) *DataServiceSQLite {
	return &DataServiceSQLite{
		repo: repo,
	}
}

func (ds *DataServiceSQLite) Create(data *models.Student, ctx context.Context) error {

	if err := ds.ValidateData(data); err != nil {
		return DataError{Message: "InvalMockDataServiceSuccessfulid data."}
	}
	return ds.repo.Create(data, ctx)
}

func (ds *DataServiceSQLite) ReadOne(id int, ctx context.Context) (*models.Student, error) {

	data, err := ds.repo.ReadOne(id, ctx)
	if err != nil {
		return nil, err
	}

	_ = data

	// Tehdään datalle jotain, päätellään datasta jotain!!!
	// Tämä ohjaa toimintaa älykkäästi, esim. jos data on tietynlaista, niin tehdään jotain

	return data, nil
}

func (ds *DataServiceSQLite) Update(data *models.Student, ctx context.Context) (int64, error) {

	if err := ds.ValidateData(data); err != nil {
		return 0, DataError{Message: "Invalid data: " + err.Error()}
	}
	return ds.repo.Update(data, ctx)
}

func (ds *DataServiceSQLite) Delete(data *models.Student, ctx context.Context) (int64, error) {
	return ds.repo.Delete(data, ctx)
}

func (ds *DataServiceSQLite) ValidateData(data *models.Student) error {
	var errMsg string
	if data.StudentName == "" || len(data.StudentName) > 50 {
		errMsg += "DeviceID is required and must be less than 50 characters. "
	}

	if errMsg != "" {
		return DataError{Message: errMsg}
	}
	return nil
}
