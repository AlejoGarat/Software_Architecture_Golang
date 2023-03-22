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

type DepartmentFactory struct {
	mongoCli *mongo.Client
	database string
	helpers  *helpers.Helpers
}

func NewDepartmentFactory(mongoCli *mongo.Client, database string, helpers *helpers.Helpers) *DepartmentFactory {
	return &DepartmentFactory{mongoCli: mongoCli, database: database, helpers: helpers}
}

func (DepartmentFactory *DepartmentFactory) GetDepartmentController() *controllers.DepartmentController {
	DepartmentRepository := InitializeDepartmentRepository(DepartmentFactory.mongoCli, DepartmentFactory.database)
	DepartmentUseCase := InitializeDepartmentUseCase(DepartmentRepository, DepartmentFactory.helpers)
	DepartmentController := InitializeDepartmentController(DepartmentUseCase)
	return DepartmentController
}

func InitializeDepartmentRepository(mongoCli *mongo.Client, database string) idataaccess.DepartmentRepository {
	DepartmentRepo := dataaccess.NewDepartmentMongoRepo(mongoCli, database)
	var DepartmentRepository idataaccess.DepartmentRepository = DepartmentRepo
	return DepartmentRepository
}

func InitializeDepartmentUseCase(DepartmentRepository idataaccess.DepartmentRepository, helpers *helpers.Helpers) *usecases.DepartmentUseCase {
	return usecases.NewDepartmentUseCase(DepartmentRepository, *helpers)
}

func InitializeDepartmentController(DepartmentUseCase iusecases.DepartmentUseCase) *controllers.DepartmentController {
	return controllers.NewDepartmentController(DepartmentUseCase)
}
