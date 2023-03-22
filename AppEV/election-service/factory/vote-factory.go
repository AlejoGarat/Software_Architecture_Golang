package factory

import (
	"election-service/api/controllers"
	"election-service/dataaccess"
	idataaccess "election-service/dataaccess/interfaces"
	"election-service/helpers"
	"election-service/usecases"
	iusecases "election-service/usecases/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
)

type VoteFactory struct {
	mongoCli *mongo.Client
	database string
	helpers  *helpers.Helpers
}

func NewVoteFactory(mongoCli *mongo.Client, database string, helpers *helpers.Helpers) *VoteFactory {
	return &VoteFactory{mongoCli: mongoCli, database: database, helpers: helpers}
}

func (VoteFactory *VoteFactory) GetVoteController() *controllers.VoteController {
	VoteRepository := InitializeVoteRepository(VoteFactory.mongoCli, VoteFactory.database)
	VoteUseCase := InitializeVoteUseCase(VoteRepository, VoteFactory.helpers)
	VoteController := InitializeVoteController(VoteUseCase)
	return VoteController
}

func InitializeVoteRepository(mongoCli *mongo.Client, database string) idataaccess.VoteRepository {
	VoteRepo := dataaccess.NewVoteMongoRepo(mongoCli, database)
	var VoteRepository idataaccess.VoteRepository = VoteRepo
	return VoteRepository
}

func InitializeVoteUseCase(VoteRepository idataaccess.VoteRepository, helpers *helpers.Helpers) *usecases.VoteUseCase {
	return usecases.NewVoteUseCase(VoteRepository, *helpers)
}

func InitializeVoteController(VoteUseCase iusecases.VoteUseCase) *controllers.VoteController {
	return controllers.NewVoteController(VoteUseCase)
}
