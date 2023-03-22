package main

import (
	"auth/factory"
	"auth/models/write"
	"auth/rabbit/workers"
	"context"
	"log"
	"os"

	"auth/api/configs"
	"auth/api/middlewares"
	"auth/api/routes"
	"auth/datasources"
	"auth/helpers"
	"auth/rabbit"
	irabbit "auth/rabbit/interfaces"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	middlewares.FiberMiddleware(app)

	openENV()

	usersDB := os.Getenv("USERS_DATABASE")

	rabbitWorker, err := getRabbitWorker()

	if err != nil {
		log.Fatalf(err.Error())
	}

	helpers := getHelpers(rabbitWorker)

	var log write.LoggingModel

	mongoAddress := os.Getenv("MONGO_CLIENT")

	mongoClient, err := getMongoClient(mongoAddress)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Mongo Connection", Actor: "User", Description: err.Error()}
		helpers.LogHelper.SendLog(log)
	}

	defer mongoClient.Disconnect(context.TODO())

	userFactory := factory.NewUserFactory(mongoClient, usersDB, *helpers)

	userController := userFactory.GetUserController()

	routes.PublicRoutes(app, userController)

	port := os.Getenv("PORT")

	app.Listen(port)
}

func openENV() {
	dotenvErr := godotenv.Load("../.env")

	if dotenvErr != nil {
		log.Fatalf("Error loading .env file")
	}
}

func getMongoClient(mongoAddress string) (*mongo.Client, error) {
	return datasources.NewMongoDataSource(mongoAddress)
}

func getRabbitWorker() (workers.Worker, error) {
	return rabbit.NewRabbitConnection()
}

func getHelpers(rabbitWorker workers.Worker) *helpers.Helpers {
	var loggerCommunication irabbit.RabbitCommunication = rabbit.NewLogRabbitCommunication(rabbitWorker)

	loggerHelper := helpers.NewLogHelper(loggerCommunication)

	helpers := helpers.Helpers{LogHelper: *loggerHelper}

	return &helpers
}
