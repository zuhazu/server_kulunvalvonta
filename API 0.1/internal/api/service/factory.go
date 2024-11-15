package service

import (
	"context"
	"goapi/internal/api/repository/DAL"
	"goapi/internal/api/repository/DAL/SQLite"
	service "goapi/internal/api/service/data"
	room_service "goapi/internal/api/service/room"
	"log"
)

type DataServiceType int
type RoomServiceType int

const (
	SQLiteDataService DataServiceType = iota
	SQLiteRoomService RoomServiceType = iota
)

type ServiceFactory struct {
	db     DAL.SQLDatabase
	logger *log.Logger
	ctx    context.Context
}

// * Factory for creating data service *
func NewServiceFactory(db DAL.SQLDatabase, logger *log.Logger, ctx context.Context) *ServiceFactory {
	return &ServiceFactory{
		db:     db,
		logger: logger,
		ctx:    ctx,
	}
}

func (sf *ServiceFactory) CreateDataService(serviceType DataServiceType) (*service.DataServiceSQLite, error) {

	switch serviceType {

	case SQLiteDataService:
		repo, err := SQLite.NewDataRepository(sf.db, sf.ctx)
		if err != nil {
			return nil, err
		}
		ds := service.NewDataServiceSQLite(repo)
		return ds, nil
	default:
		return nil, service.DataError{Message: "Invalid data service type."}
	}
}

// Lisää
func (sf *ServiceFactory) CreateRoomService(serviceType RoomServiceType) (*room_service.RoomServiceSQLite, error) {
	switch serviceType {

	case SQLiteRoomService:
		repo, err := SQLite.NewRoomRepository(sf.db, sf.ctx)
		if err != nil {
			return nil, err
		}
		rs := room_service.NewRoomServiceSQLite(repo)
		return rs, nil

	default:
		return nil, room_service.RoomError{Message: "Invalid room service type."}
	}
}
