package factory

import (
	"analytics/api/controllers"
	"analytics/dataaccess"
	idataaccess "analytics/dataaccess/interfaces"
	"analytics/helpers"
	"analytics/usecases"
	iusecases "analytics/usecases/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
)

type ScheduleFactory struct {
	mongoCli *mongo.Client
	database string
	helpers  *helpers.Helpers
}

func NewScheduleFactory(mongoCli *mongo.Client, database string, helpers *helpers.Helpers) *ScheduleFactory {
	return &ScheduleFactory{mongoCli: mongoCli, database: database, helpers: helpers}
}

func (ScheduleFactory *ScheduleFactory) GetScheduleController() *controllers.ScheduleController {
	ScheduleRepository := InitializeScheduleRepository(ScheduleFactory.mongoCli, ScheduleFactory.database)
	ScheduleUseCase := InitializeScheduleUseCase(ScheduleRepository, ScheduleFactory.helpers)
	ScheduleController := InitializeScheduleController(ScheduleUseCase)
	return ScheduleController
}

func InitializeScheduleRepository(mongoCli *mongo.Client, database string) idataaccess.ScheduleRepository {
	ScheduleRepo := dataaccess.NewScheduleMongoRepo(mongoCli, database)
	var ScheduleRepository idataaccess.ScheduleRepository = ScheduleRepo
	return ScheduleRepository
}

func InitializeScheduleUseCase(ScheduleRepository idataaccess.ScheduleRepository, helpers *helpers.Helpers) *usecases.ScheduleUseCase {
	return usecases.NewScheduleUseCase(ScheduleRepository, *helpers)
}

func InitializeScheduleController(ScheduleUseCase iusecases.ScheduleUseCase) *controllers.ScheduleController {
	return controllers.NewScheduleController(ScheduleUseCase)
}
