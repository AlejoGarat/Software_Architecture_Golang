package main

import (
	"context"
	"fmt"
	"log"
	"monitoring-service/dataaccess"
	idataaccess "monitoring-service/dataaccess/interfaces"
	"monitoring-service/datasources"
	"monitoring-service/helpers"
	"monitoring-service/models"
	"monitoring-service/rabbit"
	irabbit "monitoring-service/rabbit/interfaces"
	"monitoring-service/rabbit/workers"
	"monitoring-service/usecases"
	iusecases "monitoring-service/usecases/interfaces"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var log models.LoggingModel

	rabbitWorker := getRabbitWorker()

	var loggerCommunication irabbit.RabbitCommunication = rabbit.NewLogRabbitCommunication(rabbitWorker)

	var mailCommunication irabbit.RabbitCommunication = rabbit.NewMailRabbitCommunication(rabbitWorker)

	logHelper := helpers.NewLogHelper(loggerCommunication)

	mailHelper := helpers.NewMailHelper(mailCommunication)

	openENV()

	mongoAddress := os.Getenv("MONGO_CLIENT")

	mongoClient := getMongoClient(mongoAddress)

	defer mongoClient.Disconnect(context.TODO())

	redisAddress := os.Getenv("REDIS_CLIENT")

	redisClient := getRedisClient(redisAddress)

	defer redisClient.Close()

	electionDB := os.Getenv("ELECTION_DATABASE")

	configRepo := dataaccess.NewConfigurationRepository(mongoClient, redisClient, electionDB)

	var configRepository idataaccess.ConfigurationRepository = configRepo

	var configUseCase iusecases.ConfigurationUseCase = usecases.NewConfigurationUseCase(configRepository)

	for {
		time.Sleep(time.Second * 5)
		messages, err := configUseCase.AnalyzeValues(rabbitWorker)

		fmt.Println(len(messages))

		if err != nil {
			log = models.LoggingModel{Type: "Error", Operation: "Analyze expected values", Actor: "N/A", Description: err.Error()}
			logHelper.SendLog(log)
		}

		for _, message := range messages {
			mailHelper.SendMail(message)
		}
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
	redisClient, err := datasources.NewRedisDataSource(redisAddress)

	if err != nil {
		panic(err)
	}

	return redisClient
}

func getRabbitWorker() workers.Worker {
	rabbitWorker, err := rabbit.NewRabbitConnection()

	if err != nil {
		panic(err)
	}

	return rabbitWorker
}
