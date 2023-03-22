package factory

import (
	"election-service/adapters"
	"election-service/api/controllers"
	"election-service/dataaccess"
	idataaccess "election-service/dataaccess/interfaces"
	"election-service/helpers"
	"election-service/memorydataaccess"
	imemorydataaccess "election-service/memorydataaccess/interfaces"
	"election-service/models/write"
	"election-service/pipe"
	"election-service/usecases"
	iusecases "election-service/usecases/interfaces"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type ElectionFactory struct {
	mongoCli *mongo.Client
	redisCli *redis.Client
	database string
	helpers  *helpers.Helpers
}

func NewElectionFactory(mongoCli *mongo.Client, redisCli *redis.Client, database string, helpers *helpers.Helpers) *ElectionFactory {
	return &ElectionFactory{mongoCli: mongoCli, redisCli: redisCli, database: database, helpers: helpers}
}

func (ElectionFactory *ElectionFactory) GetElectionController() *controllers.ElectionController {
	ElectionRepository, IMElectionRepository := InitializeElectionsRepository(ElectionFactory.mongoCli, ElectionFactory.redisCli, ElectionFactory.database)
	ElectionUseCase := InitializeElectionUseCase(ElectionRepository, IMElectionRepository, ElectionFactory.helpers)
	ElectionController := InitializeElectionController(ElectionUseCase)
	return ElectionController
}

func InitializeElectionsRepository(mongoCli *mongo.Client, redisCli *redis.Client, database string) (idataaccess.ElectionRepository, imemorydataaccess.ElectionMemoryRepository) {
	ElectionRepo := dataaccess.NewElectionMongoRepo(mongoCli, database)
	MemoryRepo := memorydataaccess.NewRedisElectionImp(redisCli)
	var ElectionRepository idataaccess.ElectionRepository = ElectionRepo
	var MemoryElectionRepository imemorydataaccess.ElectionMemoryRepository = MemoryRepo
	return ElectionRepository, MemoryElectionRepository
}

func InitializeElectionUseCase(ElectionRepository idataaccess.ElectionRepository, ElectionInMemoryRepository imemorydataaccess.ElectionMemoryRepository, helpers *helpers.Helpers) *usecases.ElectionUsecase {
	startElectionPipeline, endElectionPipeline := createPipelines(ElectionInMemoryRepository, ElectionRepository, helpers)
	jsonAdapter := adapters.ElectionJsonAdapter{}
	return usecases.NewElectionUseCase(ElectionRepository, jsonAdapter, ElectionInMemoryRepository, startElectionPipeline, endElectionPipeline, *helpers)
}

func InitializeElectionController(ElectionUseCase iusecases.ElectionUsecase) *controllers.ElectionController {
	return controllers.NewElectionController(ElectionUseCase)
}

func createPipelines(memoryRepo imemorydataaccess.ElectionMemoryRepository, repository idataaccess.ElectionRepository, helpers *helpers.Helpers) (*pipe.Pipeline, *pipe.Pipeline) {
	var log write.LoggingModel

	startElectionPipeline := pipe.NewPipeline(repository)

	startFilters, err := memoryRepo.GetStartElectionFilters()

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Mongo Connection", Actor: "Add Election Start Filters", Description: err.Error()}
		helpers.LogHelper.SendLog(log)
	}

	startElectionPipeline.AddFiltersToExecute(startFilters.Filters)

	closeElectionPipeline := pipe.NewPipeline(repository)

	endFilters, err := memoryRepo.GetCloseElectionFilters()

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Mongo Connection", Actor: "Add Election End Filters", Description: err.Error()}
		helpers.LogHelper.SendLog(log)
	}

	closeElectionPipeline.AddFiltersToExecute(endFilters.Filters)

	return startElectionPipeline, closeElectionPipeline
}
