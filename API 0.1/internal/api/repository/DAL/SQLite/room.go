package SQLite

import (
	"context"
	"database/sql"
	"goapi/internal/api/repository/DAL"
	"goapi/internal/api/repository/models"
)

type RoomRepository struct {
	sqlDB *sql.DB
	createStmt,
	readPersonsByRoomIdStmt,
	readStmt *sql.Stmt
	ctx context.Context
}

func NewRoomRepository(sqlDB DAL.SQLDatabase, ctx context.Context) (models.RoomRepository, error) {

	repo := &RoomRepository{
		sqlDB: sqlDB.Connection(),
		ctx:   ctx,
	}

	// Create the room table if it doesn't exist
	if _, err := repo.sqlDB.Exec(`CREATE TABLE IF NOT EXISTS room (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		room_id VARCHAR(50) NOT NULL,
		room_name VARCHAR(50)
	);`); err != nil {
		repo.sqlDB.Close()
		return nil, err
	}

	// Prepare SQL statements
	createStmt, err := repo.sqlDB.Prepare(`INSERT INTO room (room_id, room_name) VALUES (?, ?)`)
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.createStmt = createStmt

	readStmt, err := repo.sqlDB.Prepare("SELECT id, room_id, room_name FROM room WHERE id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readStmt = readStmt

	readPersonsByRoomIdStmt, err := repo.sqlDB.Prepare("SELECT person_name FROM person WHERE room_id = ?")
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readPersonsByRoomIdStmt = readPersonsByRoomIdStmt

	// Ensure that resources are cleaned up after context is done
	go CloseRoom(ctx, repo)

	return repo, nil
}

func CloseRoom(ctx context.Context, r *RoomRepository) {
	<-ctx.Done()
	r.createStmt.Close()
	r.readStmt.Close()
	r.readPersonsByRoomIdStmt.Close()
	r.sqlDB.Close()
}

func (r *RoomRepository) CreateRoom(room *models.Room, ctx context.Context) error {
	res, err := r.createStmt.ExecContext(ctx, room.RoomID, room.RoomName)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	room.ID = int(id)
	return nil
}

func (r *RoomRepository) ReadOneRoom(id int, ctx context.Context) (*models.Room, error) {
	row := r.readStmt.QueryRowContext(ctx, id)
	var room models.Room
	err := row.Scan(&room.ID, &room.RoomID, &room.RoomName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) GetPersonsByRoomID(room_id string, ctx context.Context) ([]*models.Person, error) {

	rows, err := r.readPersonsByRoomIdStmt.QueryContext(ctx, room_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []*models.Person
	for rows.Next() {
		var d models.Person
		err := rows.Scan(&d.PersonName)
		if err != nil {
			return nil, err
		}
		data = append(data, &d)
	}
	return data, nil
}
