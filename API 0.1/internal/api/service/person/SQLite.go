package person

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

// Luodaan person
func (ds *PersonServiceSQLite) CreatePerson(data *models.Person, ctx context.Context) error {

	if err := ds.ValidateData(data); err != nil {
		return PersonError{Message: err.Error()}
	}
	return ds.repo.CreatePerson(data, ctx)
}

// Haetaan yksi person
func (ds *PersonServiceSQLite) ReadOnePerson(id int, ctx context.Context) (*models.Person, error) {

	data, err := ds.repo.ReadOnePerson(id, ctx)
	if err != nil {
		return nil, err
	}

	_ = data

	return data, nil
}

// Päivitetään person
func (ds *PersonServiceSQLite) UpdatePerson(data *models.Person, ctx context.Context) (int64, error) {

	if err := ds.ValidateData(data); err != nil {
		return 0, PersonError{Message: "Invalid data: " + err.Error()}
	}
	return ds.repo.UpdatePerson(data, ctx)
}

// Poistetaan person
func (ds *PersonServiceSQLite) DeletePerson(data *models.Person, ctx context.Context) (int64, error) {
	return ds.repo.DeletePerson(data, ctx)
}

// Päivitetään roomID tagID:n perusteella
func (ps *PersonServiceSQLite) UpdateRoomIDByTagID(tagID, newRoomID string, ctx context.Context) (string, error) {
	// Kutsutaan repositoryn metodia
	message, err := ps.repo.UpdateRoomIDByTagID(tagID, newRoomID, ctx)
	if err != nil {
		return "Access denied", err
	}

	// Palautetaan palvelun tulos
	return message, nil
}

// Validointi
func (ds *PersonServiceSQLite) ValidateData(data *models.Person) error {
	var errMsg string
	if data.PersonName == "" || len(data.PersonName) > 50 || len(data.PersonName) < 4 {
		errMsg += "Invalid name."
	}
	if errMsg != "" {
		return PersonError{Message: errMsg}
	}
	return nil
}

// Haetaan reposta lista henkilöistä room id:n perusteella
func (ps *PersonServiceSQLite) ReadPersonsByRoomId(roomId string, ctx context.Context) ([]*models.Person, error) {
	data, err := ps.repo.ReadPersonsByRoomId(roomId, ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
