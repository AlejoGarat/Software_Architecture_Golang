package factory

import (
	"consultant-service/api/controllers"
	"consultant-service/dataaccess"
	idataaccess "consultant-service/dataaccess/interfaces"
	"consultant-service/helpers"
	"consultant-service/usecases"
	iusecases "consultant-service/usecases/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
)

type CircuitFactory struct {
	mongoCli *mongo.Client
	database string
	helpers  *helpers.Helpers
}

func NewCircuitFactory(mongoCli *mongo.Client, database string, helpers *helpers.Helpers) *CircuitFactory {
	return &CircuitFactory{mongoCli: mongoCli, database: database, helpers: helpers}
}

func (circuitFactory *CircuitFactory) GetCircuitController() *controllers.CircuitController {
	circuitRepository := InitializeCircuitRepository(circuitFactory.mongoCli, circuitFactory.database)
	circuitUseCase := InitializeCircuitUseCase(circuitRepository, circuitFactory.helpers)
	circuitController := InitializeCircuitController(circuitUseCase)
	return circuitController
}

func InitializeCircuitRepository(mongoCli *mongo.Client, database string) idataaccess.CircuitRepository {
	circuitRepo := dataaccess.NewCircuitMongoRepo(mongoCli, database)
	var circuitRepository idataaccess.CircuitRepository = circuitRepo
	return circuitRepository
}

func InitializeCircuitUseCase(circuitRepository idataaccess.CircuitRepository, helpers *helpers.Helpers) *usecases.CircuitUseCase {
	return usecases.NewCircuitUseCase(circuitRepository, *helpers)
}

func InitializeCircuitController(circuitUseCase iusecases.CircuitUseCase) *controllers.CircuitController {
	return controllers.NewCircuitController(circuitUseCase)
}
