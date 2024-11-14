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
	updateStmt,
	deleteStmt,
	updateRoomIDStmt,
	readRoomIDStmt,
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

	deleteStmt, err := repo.sqlDB.Prepare("DELETE FROM person WHERE id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.deleteStmt = deleteStmt

	//Ei haluta tässä päivittää roomID:tä
	updateStmt, err := repo.sqlDB.Prepare("UPDATE person SET person_id = ?, person_name = ? WHERE id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.updateStmt = updateStmt

	updateRoomIDStmt, err := repo.sqlDB.Prepare("UPDATE person SET room_id = ? WHERE person_id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.updateRoomIDStmt = updateRoomIDStmt

	readRoomIDStmt, err := repo.sqlDB.Prepare("SELECT room_id FROM person WHERE person_id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readRoomIDStmt = readRoomIDStmt

	// Ensure that resources are cleaned up after context is done
	go ClosePerson(ctx, repo)

	return repo, nil
}

func ClosePerson(ctx context.Context, r *PersonRepository) {
	<-ctx.Done()
	r.createStmt.Close()
	r.readStmt.Close()
	r.updateStmt.Close()
	r.updateRoomIDStmt.Close()
	r.readRoomIDStmt.Close()
	r.deleteStmt.Close()
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

func (r *PersonRepository) UpdatePerson(data *models.Person, ctx context.Context) (int64, error) {
	res, err := r.updateStmt.ExecContext(ctx, data.PersonID, data.PersonName, data.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (r *PersonRepository) DeletePerson(data *models.Person, ctx context.Context) (int64, error) {
	res, err := r.deleteStmt.ExecContext(ctx, data.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (r *PersonRepository) UpdateRoomIDByTagID(personID, tagID, newRoomID string, ctx context.Context) (string, error) {
	var currentRoomID string
	row := r.readRoomIDStmt.QueryRowContext(ctx, personID)
	if err := row.Scan(&currentRoomID); err != nil {
		if err == sql.ErrNoRows {
			// Ei löydy henkilöä
			return "epäonnistui", err
		}
		return "epäonnistui", err
	}

	if currentRoomID != "-1" {
		_, err := r.updateRoomIDStmt.ExecContext(ctx, "-1", personID)
		if err != nil {
			return "epäonnistui", err
		}
		return "kirjauduttu ulos", nil
	}

	if currentRoomID == "-1" {
		_, err := r.updateRoomIDStmt.ExecContext(ctx, tagID, personID)
		if err != nil {
			return "epäonnistui", err
		}
		return "kirjauduttu sisään", nil
	}

	return "epäonnistui", nil
}
