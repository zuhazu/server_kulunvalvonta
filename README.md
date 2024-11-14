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
