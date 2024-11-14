# server_kulunvalvonta
Älykkäät laitteet serveri puoli

# luo uusi entiteetti ohjeet
1. Luo model -> internal/api/repository/models : person.go
```go
package models

import "context"

type Person struct {
	ID         int    `json:"id"`
	PersonID   int    `json:"person_id"`
	PersonName string `json:"person_name"`
	RoomID     int    `json:"room_id"`
}

type PersonRepository interface {
	CreatePerson(Data *Person, ctx context.Context) error
	ReadOnePerson(id int, ctx context.Context) (*Person, error)
}
```

2. Luo repository -> internal/api/repository/DAL/SQLite : person.go
```go
package SQLite

import (
	"context"
	"database/sql"
	"goapi/internal/api/repository/DAL"
	"goapi/internal/api/repository/models"
)

type PersonRepository struct {
	sqlDB *sql.DB
	createStmt,
	readStmt *sql.Stmt
	ctx context.Context
}

func NewPersonRepository(sqlDB DAL.SQLDatabase, ctx context.Context) (models.PersonRepository, error) {

	repo := &PersonRepository{
		sqlDB: sqlDB.Connection(),
		ctx:   ctx,
	}

	// Create the person table if it doesn't exist
	if _, err := repo.sqlDB.Exec(`CREATE TABLE IF NOT EXISTS person (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		person_id VARCHAR(50) NOT NULL,
		person_name VARCHAR(50),
		room_id VARCHAR(50)
	);`); err != nil {
		repo.sqlDB.Close()
		return nil, err
	}

	// Prepare SQL statements
	createStmt, err := repo.sqlDB.Prepare(`INSERT INTO person (person_id, person_name, room_id) VALUES (?, ?, ?)`)
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.createStmt = createStmt

	readStmt, err := repo.sqlDB.Prepare("SELECT id, person_id, person_name, room_id FROM person WHERE id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readStmt = readStmt

	// Ensure that resources are cleaned up after context is done
	go ClosePerson(ctx, repo)

	return repo, nil
}

func ClosePerson(ctx context.Context, r *PersonRepository) {
	<-ctx.Done()
	r.createStmt.Close()
	r.readStmt.Close()
	r.sqlDB.Close()
}

func (r *PersonRepository) CreatePerson(person *models.Person, ctx context.Context) error {
	res, err := r.createStmt.ExecContext(ctx, person.PersonID, person.PersonName, person.RoomID)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	person.ID = int(id)
	return nil
}

func (r *PersonRepository) ReadOnePerson(id int, ctx context.Context) (*models.Person, error) {
	row := r.readStmt.QueryRowContext(ctx, id)
	var person models.Person
	err := row.Scan(&person.ID, &person.PersonID, &person.PersonName, &person.RoomID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &person, nil
}
```
3.1 Luo service -> internal/api/service/[palvelun nimi] : common.go
```go
package data

import (
	"context"
	"goapi/internal/api/repository/models"
)

type PersonService interface {
	CreatePerson(data *models.Person, ctx context.Context) error
	ReadOnePerson(id int, ctx context.Context) (*models.Person, error)
	ValidateData(data *models.Person) error
}

type PersonError struct {
	Message string
}

func (de PersonError) Error() string {
	return de.Message
}
```
3.2 Luo internal/api/service/[palvelun nimi] : SQLite.go
```go
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

func (ds *PersonServiceSQLite) ValidateData(data *models.Person) error {
	var errMsg string
	if errMsg != "" {
		return PersonError{Message: errMsg}
	}
	return nil
}
```
3.3 Luo internal/api/service/[palvelun nimi] : mocks.go
4 Lisää koodit internal/api/service : factory.go
```go
type DataServiceType int
type PersonServiceType int // Lisää

const (
	SQLiteDataService   DataServiceType   = iota
	SQLitePersonService PersonServiceType = iota // Lisää
)
```
```go
// Lisää
func (sf *ServiceFactory) CreatePersonService(serviceType PersonServiceType) (*person_service.PersonServiceSQLite, error) {
	switch serviceType {

	case SQLitePersonService:
		repo, err := SQLite.NewPersonRepository(sf.db, sf.ctx)
		if err != nil {
			return nil, err
		}
		ps := person_service.NewPersonServiceSQLite(repo)
		return ps, nil

	default:
		return nil, person_service.PersonError{Message: "Invalid person service type."}
	}
}
```
5. Lisää koodit setupDataHandlers funktioon : server.go
```go
//lisää
	personService, err := sf.CreatePersonService(service.SQLitePersonService)
	if err != nil {
		return err
	}
```
```go
//Lisää handlerit
	mux.HandleFunc("POST /person", func(w http.ResponseWriter, r *http.Request) {
		data.PostPersonHandler(w, r, logger, personService)
	})
```
6. Luo handlerit internal/api/handlers/ : esim: postperson.go
```
package data

import (
	"context"
	"encoding/json"
	"goapi/internal/api/repository/models"
	service "goapi/internal/api/service/person"
	"log"
	"net/http"
	"time"
)

// * User sends a POST request to /data with a JSON payload in the request body *
// * curl -X POST http://127.0.0.1:8080/data -i -u admin:password -H "Content-Type: application/json" -d '{"device_id": "device1", "device_name": "device1", "value": 1.0, "type": "type1", "date_time": "2021-01-01T00:00:00Z", "description": "description1"}'
func PostPersonHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.PersonService) {
	var data models.Person

	// * Decode the JSON payload from the request body into the data struct
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {

		// * This is a User Error: format of body is invalid, response in JSON and with a 400 status code
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid request data. Please check your input."}`))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// * Try to create the data in the database
	if err := ds.CreatePerson(&data, ctx); err != nil {
		switch err.(type) {
		case service.PersonError:
			// * If the error is a DataError, handle it as a client error
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "` + err.Error() + `"}`))
			return
		default:
			// * If it is not a DataError, handle it as a server error
			logger.Println("Error creating data:", err, data)
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
			return
		}
	}

	// * Return the data to the user as JSON with a 201 Created status code
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Println("Error encoding data:", err, data)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}
}
```
