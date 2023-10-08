package registry

import (
	"github.com/xkurozaru/plant-diagnosis/controller/application"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/repository"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/service"
	"github.com/xkurozaru/plant-diagnosis/controller/infrastructure/database"
	"github.com/xkurozaru/plant-diagnosis/controller/interface/handler"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

func NewRegistry(db *gorm.DB) *registry {
	return &registry{db: db}
}

// Repository //

func (r registry) NewAuthenticationRepository() repository.AuthenticationRepository {
	return database.NewAuthenticationDatabase(r.db)
}

func (r registry) NewUserRepository() repository.UserRepository {
	return database.NewUserDatabase(r.db)
}

func (r registry) NewPredictionModelRepository() repository.PredictionModelRepository {
	return database.NewPredictionModelDatabase(r.db)
}

func (r registry) NewPredictionResultRepository() repository.PredictionResultRepository {
	return database.NewPredictionResultDatabase(r.db)
}

// Domain Service //

func (r registry) NewPredictionService() service.PredictionService {
	return service.NewPredictionService(
		r.NewPredictionResultRepository(),
	)
}

// Application Service //

func (r registry) NewAccountApplicationService() application.AccountApplicationService {
	return application.NewAccountApplicationService(
		r.NewAuthenticationRepository(),
		r.NewUserRepository(),
	)
}

func (r registry) NewPredictionApplicationService() application.PredictionApplicationService {
	return application.NewPredictionApplicationService(
		r.NewPredictionModelRepository(),
		r.NewPredictionResultRepository(),
		r.NewUserRepository(),
		r.NewPredictionService(),
	)
}

// Handler //

func (r registry) NewAccountHandler() handler.AccountHandler {
	return handler.NewAccountHandler(
		r.NewAccountApplicationService(),
	)
}

func (r registry) NewPredictionHandler() handler.PredictionHandler {
	return handler.NewPredictionHandler(
		r.NewPredictionApplicationService(),
	)
}
