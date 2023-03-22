package factory

import (
	"votation-service/api/controllers"
	"votation-service/dataaccess"
	idataaccess "votation-service/dataaccess/interfaces"
	"votation-service/helpers"
	"votation-service/messenger"
	"votation-service/usecases"
	iusecases "votation-service/usecases/interfaces"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type VoteFactory struct {
	mongoCli *mongo.Client
	redisCli *redis.Client
	database string
	helpers  helpers.Helpers
}

func NewVoteFactory(mongoCli *mongo.Client, redisCli *redis.Client, database string, helpers helpers.Helpers) *VoteFactory {
	return &VoteFactory{mongoCli: mongoCli, redisCli: redisCli, database: database, helpers: helpers}
}

func (voteFactory *VoteFactory) GetVoteController() (*controllers.VoteController, *idataaccess.VoteRepository) {
	voteRepository := InitializeVoteRepository(voteFactory.mongoCli, voteFactory.redisCli, voteFactory.database)
	voteUseCase := InitializeVoteUseCase(voteRepository, voteFactory.helpers)
	voteController := InitializeVoteController(voteUseCase)
	return voteController, &voteRepository
}

func InitializeVoteRepository(mongoCli *mongo.Client, redisCli *redis.Client, database string) idataaccess.VoteRepository {
	voteRepo := dataaccess.NewVotesRepo(mongoCli, redisCli, database)
	var voteRepository idataaccess.VoteRepository = voteRepo
	return voteRepository
}

func InitializeVoteUseCase(voteRepository idataaccess.VoteRepository, helpers helpers.Helpers) *usecases.VoteUseCase {
	smsMessenger := messenger.NewSms(helpers.MessageHelper)
	return usecases.NewVoteUseCase(voteRepository, smsMessenger, helpers)
}

func InitializeVoteController(voteUseCase iusecases.VoteUseCase) *controllers.VoteController {
	return controllers.NewVoteController(voteUseCase)
}
