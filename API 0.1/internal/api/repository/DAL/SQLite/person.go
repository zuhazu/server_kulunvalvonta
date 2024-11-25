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
	readRoomStmt,
	updateRoomIDStmt,
	readRoomIDStmt,
	readPersonsByRoomIdStmt,
	readStmt *sql.Stmt
	ctx context.Context
}

func NewPersonRepository(sqlDB DAL.SQLDatabase, ctx context.Context) (models.PersonRepository, error) {

	repo := &PersonRepository{
		sqlDB: sqlDB.Connection(),
		ctx:   ctx,
	}

	// Luodaan taulu jos sitä ei ole olemassa
	if _, err := repo.sqlDB.Exec(`CREATE TABLE IF NOT EXISTS person (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		person_id VARCHAR(50) NOT NULL,
		person_name VARCHAR(50),
		tag_id VARCHAR(50),
		room_id VARCHAR(50)
	);`); err != nil {
		repo.sqlDB.Close()
		return nil, err
	}

	// Prepare SQL statements
	createStmt, err := repo.sqlDB.Prepare(`INSERT INTO person (person_id, person_name, room_id, tag_id) VALUES (?, ?, ?, ?)`)
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.createStmt = createStmt

	readStmt, err := repo.sqlDB.Prepare("SELECT id, person_id, person_name, room_id, tag_id FROM person WHERE id = ?")
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
	updateStmt, err := repo.sqlDB.Prepare("UPDATE person SET person_id = ?, person_name = ?, tag_id = ? WHERE id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.updateStmt = updateStmt

	// Haetaan henkilöt, joiden room_id on tietty
	readPersonsByRoomIdStmt, err := repo.sqlDB.Prepare("SELECT id, room_id, person_id, person_name, tag_id FROM person WHERE room_id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readPersonsByRoomIdStmt = readPersonsByRoomIdStmt

	updateRoomIDStmt, err := repo.sqlDB.Prepare("UPDATE person SET room_id = ? WHERE tag_id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.updateRoomIDStmt = updateRoomIDStmt

	readRoomIDStmt, err := repo.sqlDB.Prepare("SELECT room_id FROM person WHERE tag_id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readRoomIDStmt = readRoomIDStmt

	readRoomStmt, err := repo.sqlDB.Prepare("SELECT id, room_id, room_name FROM room WHERE room_id = ? LIMIT 1")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readRoomStmt = readRoomStmt

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
	r.readRoomStmt.Close()
	r.readPersonsByRoomIdStmt.Close()
	r.sqlDB.Close()
}

func (r *PersonRepository) CreatePerson(person *models.Person, ctx context.Context) error {
	res, err := r.createStmt.ExecContext(ctx, person.PersonID, person.PersonName, person.RoomID, person.TagID)
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
	err := row.Scan(&person.ID, &person.PersonID, &person.PersonName, &person.RoomID, &person.TagID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &person, nil
}

func (r *PersonRepository) UpdatePerson(data *models.Person, ctx context.Context) (int64, error) {
	res, err := r.updateStmt.ExecContext(ctx, data.PersonID, data.PersonName, data.TagID, data.ID)
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

func (r *PersonRepository) ReadPersonsByRoomId(roomId string, ctx context.Context) ([]*models.Person, error) {
	// Suoritetaan kysely, joka palauttaa useita rivejä
	rows, err := r.readPersonsByRoomIdStmt.QueryContext(ctx, roomId)
	if err != nil {
		return nil, err
	}

	var persons []*models.Person

	// Käydään läpi jokainen rivi
	for rows.Next() {
		var person models.Person
		err := rows.Scan(&person.ID, &person.PersonID, &person.PersonName, &person.RoomID, &person.TagID)
		if err != nil {
			return nil, err
		}
		persons = append(persons, &person) // Lisätään person listaan
	}

	// Tarkistetaan virheet, jotka tapahtuivat rivien iteraation aikana
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return persons, nil
}

func (r *PersonRepository) UpdateRoomIDByTagID(tagID, newRoomID string, ctx context.Context) (string, error) {
	roomRow, error2 := r.readRoomStmt.QueryContext(ctx, newRoomID)
	if error2 != nil {
		return "Unexpected error", error2
	}
	if !roomRow.Next() {
		return "Room not found", nil
	}

	var currentRoomID string
	row := r.readRoomIDStmt.QueryRowContext(ctx, tagID)
	if err := row.Scan(&currentRoomID); err != nil {
		if err == sql.ErrNoRows {
			// Ei löydy henkilöä
			return "Access denied", err
		}
		return "Access denied", err
	}

	if currentRoomID != "-1" {
		_, err := r.updateRoomIDStmt.ExecContext(ctx, "-1", tagID)
		if err != nil {
			return "Access denied", err
		}
		return "Logged out", nil
	}

	if currentRoomID == "-1" {
		_, err := r.updateRoomIDStmt.ExecContext(ctx, newRoomID, tagID)
		if err != nil {
			return "Access denied", err
		}
		return "Logged in", nil
	}

	return "Access denied", nil
}
