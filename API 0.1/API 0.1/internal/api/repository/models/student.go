package models

import "context"

type Student struct {
	ID          int    `json:"id"`
	StudentID   int    `json:"student_id"`
	StudentName string `json:"student_name"`
	RoomID      int    `json:"room_id"`
}

type StudentRepository interface {
	Create(Data *Student, ctx context.Context) error
	ReadOne(id int, ctx context.Context) (*Student, error)
	// ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*Data, error)
	Update(data *Student, ctx context.Context) (int64, error)
	Delete(data *Student, ctx context.Context) (int64, error)
}
