package service

import (
	"context"
	"goapi/internal/api/repository/DAL"
	"goapi/internal/api/repository/DAL/SQLite"
	data_service "goapi/internal/api/service/data"
	person_service "goapi/internal/api/service/person"
	"log"
)

type DataServiceType int
type PersonServiceType int

const (
	SQLiteDataService   DataServiceType   = iota
	SQLitePersonService PersonServiceType = iota
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
