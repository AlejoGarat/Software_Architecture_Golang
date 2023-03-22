package main

import (
	"context"
	"log"
	"os"
	"votation-service/factory"
	"votation-service/pipe"
	"votation-service/rabbit/workers"

	"votation-service/api/configs"
	"votation-service/api/middlewares"
	"votation-service/api/routes"
	"votation-service/datasources"
	"votation-service/helpers"
	"votation-service/rabbit"
	irabbit "votation-service/rabbit/interfaces"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	middlewares.FiberMiddleware(app)

	openENV()

	mongoAddress := os.Getenv("MONGO_CLIENT")

	mongoClient := getMongoClient(mongoAddress)

	defer mongoClient.Disconnect(context.TODO())

	redisClient := getRedisClient("localhost:6379")

	defer redisClient.Close()

	votationDB := os.Getenv("VOTATION_DATABASE")

	rabbitWorker := getRabbitWorker()

	helpers := getHelpers(rabbitWorker)

	voteFactory := factory.NewVoteFactory(mongoClient, redisClient, votationDB, *helpers)

	filtersFactory := factory.NewFilterFactory(redisClient)

	voteController, voteRepository := voteFactory.GetVoteController()

	filtersRepository := filtersFactory.InitializeFilterRepository(redisClient)

	pipeline := pipe.NewPipeline(*voteRepository, filtersRepository)

	voteIssuance := middlewares.NewVoteIssuanceMiddleware(*pipeline, helpers.LogHelper)

	routes.PublicRoutes(app, voteController, voteIssuance)

	port := os.Getenv("PORT")

	app.Listen(port)
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

func getHelpers(rabbitWorker workers.Worker) *helpers.Helpers {
	var loggerCommunication irabbit.RabbitCommunication = rabbit.NewLogRabbitCommunication(rabbitWorker)
	var dataUpdateCommunication irabbit.RabbitCommunication = rabbit.NewDataUpdateRabbitCommunication(rabbitWorker)
	var messageCommunication irabbit.RabbitCommunication = rabbit.NewMessageRabbitCommunication(rabbitWorker)
	var mailCommunication irabbit.RabbitCommunication = rabbit.NewMailRabbitCommunication(rabbitWorker)

	loggerHelper := helpers.NewLogHelper(loggerCommunication)
	dataUpdateHelper := helpers.NewDataUpdateHelper(dataUpdateCommunication)
	messageHelper := helpers.NewMessageHelper(messageCommunication)
	mailHelper := helpers.NewMailHelper(mailCommunication)

	helpers := helpers.Helpers{LogHelper: *loggerHelper, DataUpdateHelper: *dataUpdateHelper, MessageHelper: *messageHelper, MailHelper: *mailHelper}

	return &helpers
}
