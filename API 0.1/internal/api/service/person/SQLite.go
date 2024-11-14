package data

import (
	"context"
	"goapi/internal/api/repository/models"
)

type PersonServiceSQLite struct {
	repo models.PersonRepository
}

func NewPersonServiceSQLite(repo models.PersonRepository) *PersonServiceSQLite {
	return &PersonServiceSQLite{
		repo: repo,
	}
}

func (ds *PersonServiceSQLite) CreatePerson(data *models.Person, ctx context.Context) error {

	if err := ds.ValidateData(data); err != nil {
		return PersonError{Message: "InvalMockDataServiceSuccessfulid data."}
	}
	return ds.repo.CreatePerson(data, ctx)
}

func (ds *PersonServiceSQLite) ReadOnePerson(id int, ctx context.Context) (*models.Person, error) {

	data, err := ds.repo.ReadOnePerson(id, ctx)
	if err != nil {
		return nil, err
	}

	_ = data

	return data, nil
}

func (ds *PersonServiceSQLite) UpdatePerson(data *models.Person, ctx context.Context) (int64, error) {

	if err := ds.ValidateData(data); err != nil {
		return 0, PersonError{Message: "Invalid data: " + err.Error()}
	}
	return ds.repo.UpdatePerson(data, ctx)
}

func (ds *PersonServiceSQLite) DeletePerson(data *models.Person, ctx context.Context) (int64, error) {
	return ds.repo.DeletePerson(data, ctx)
}

func (ps *PersonServiceSQLite) UpdateRoomIDByTagID(personID, tagID, newRoomID string, ctx context.Context) (string, error) {
	// Kutsutaan repositoryn metodia
	message, err := ps.repo.UpdateRoomIDByTagID(personID, tagID, newRoomID, ctx)
	if err != nil {
		return "epäonnistui", err
	}

	// Palautetaan palvelun tulos
	return message, nil
}
func (ds *PersonServiceSQLite) ValidateData(data *models.Person) error {
	var errMsg string
	if errMsg != "" {
		return PersonError{Message: errMsg}
	}
	return nil
}
