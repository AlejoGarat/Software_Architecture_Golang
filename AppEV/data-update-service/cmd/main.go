package main

import (
	"context"
	"data-update-service/dataaccess"
	"data-update-service/datasources"
	"data-update-service/memorydataaccess"
	imemorydataaccess "data-update-service/memorydataaccess/interfaces"
	"data-update-service/models/read"
	"data-update-service/usecases"
	iusecase "data-update-service/usecases/interfaces"
	"data-update-service/workers"
	"encoding/json"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	openENV()

	mongoAddress := os.Getenv("MONGO_CLIENT")

	mongoClient := getMongoClient(mongoAddress)

	defer mongoClient.Disconnect(context.TODO())

	redisAddress := os.Getenv("REDIS_CLIENT")

	redisClient := getRedisClient(redisAddress)

	defer redisClient.Close()

	rabbitAddress := os.Getenv("RABBIT")

	worker, err := workers.BuildRabbitWorker(rabbitAddress)

	if err != nil {
		panic(err)
	}

	votationDB := os.Getenv("ELECTORAL_DATABASE")

	voteDataRepo := dataaccess.NewVoteDataRepository(mongoClient, votationDB)

	if err != nil {
		panic(err)
	}

	var voteMemoryRepo imemorydataaccess.MemoryRepository
	voteMemoryRepo = memorydataaccess.NewRedisElectionImp(redisClient)

	var voteDataUseCase iusecase.VoteDataUseCase
	voteDataUseCase = usecases.NewVoteDataUseCase(voteDataRepo, voteMemoryRepo)

	go listenForVoteUpdates(worker, voteDataUseCase)

	listenForElectionUpdates(worker, voteDataUseCase)

	worker.Close()
}

func listenForVoteUpdates(worker workers.Worker, voteDataUseCase iusecase.VoteDataUseCase) {

	for {
		worker.Listen(100, "vote-update-queue", func(marshalledVote []byte) error {
			var vote read.Vote

			err := json.Unmarshal(marshalledVote, &vote)

			if err != nil {
				panic(err)
			}

			updateErr := voteDataUseCase.UpdateVoteData(vote)

			if updateErr != nil {
				panic(updateErr)
			}

			return nil
		})
	}
}

func listenForElectionUpdates(worker workers.Worker, voteDataUseCase iusecase.VoteDataUseCase) {

	for {
		worker.Listen(100, "update-data-queue", func(marshalledElection []byte) error {
			var election read.Election

			err := json.Unmarshal(marshalledElection, &election)

			if err != nil {
				panic(err)
			}

			updateErr := voteDataUseCase.StartUpdateCrone(election)

			if updateErr != nil {
				panic(updateErr)
			}

			return nil
		})
	}
}

func openENV() {
	dotenvErr := godotenv.Load("../.env")

	if dotenvErr != nil {
		log.Fatalf("Error loading .env file")
	}
}

func getMongoClient(mongoAddress string) *mongo.Client {
	mongoClient, err := datasources.NewMongoDataSource(mongoAddress)

	if err != nil {
		panic(err)
	}

	return mongoClient
}

func getRedisClient(redisAddress string) *redis.Client {
	redisCli, err := datasources.NewRedisDataSource(redisAddress)

	if err != nil {
		panic(err)
	}

	return redisCli

}
