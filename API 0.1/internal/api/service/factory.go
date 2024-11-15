package service

import (
	"context"
	"goapi/internal/api/repository/DAL"
	"goapi/internal/api/repository/DAL/SQLite"
	data_service "goapi/internal/api/service/data"
	person_service "goapi/internal/api/service/person"
	room_service "goapi/internal/api/service/room"
	"log"
)

type DataServiceType int
type PersonServiceType int // Lisää
type RoomServiceType int

const (
	SQLiteDataService   DataServiceType   = iota
	SQLitePersonService PersonServiceType = iota // Lisää
	SQLiteRoomService   RoomServiceType   = iota
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

func (sf *ServiceFactory) CreateDataService(serviceType DataServiceType) (*data_service.DataServiceSQLite, error) {
	switch serviceType {

	case SQLiteDataService:
		repo, err := SQLite.NewDataRepository(sf.db, sf.ctx)
		if err != nil {
			return nil, err
		}
		ds := data_service.NewDataServiceSQLite(repo)
		return ds, nil
	default:
		return nil, data_service.DataError{Message: "Invalid data service type."}
	}
}

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
