package main

import (
	"analytics/api/common"
	"analytics/api/configs"
	"analytics/api/middlewares"
	"analytics/api/routes"
	"analytics/datasources"
	"analytics/factory"
	"analytics/helpers"
	"analytics/models/write"
	"analytics/rabbit"
	irabbit "analytics/rabbit/interfaces"
	"analytics/rabbit/workers"
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	middlewares.FiberMiddleware(app)

	openENV()

	rabbitWorker := getRabbitWorker()

	helpers := getHelpers(rabbitWorker)

	var log write.LoggingModel

	mongoAddress := os.Getenv("MONGO_CLIENT")

	electoralDB := os.Getenv("ELECTORAL_DB")

	mongoClient, err := getMongoClient(mongoAddress)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Mongo Connection", Actor: "Consulting Agent", Description: err.Error()}
		helpers.LogHelper.SendLog(log)
	}

	defer mongoClient.Disconnect(context.TODO())

	circuitFactory := factory.NewCircuitFactory(mongoClient, electoralDB, &helpers)
	departmentFactory := factory.NewDepartmentFactory(mongoClient, electoralDB, &helpers)
	scheduleFactory := factory.NewScheduleFactory(mongoClient, electoralDB, &helpers)

	circuitController := circuitFactory.GetCircuitController()
	departmentController := departmentFactory.GetDepartmentController()
	scheduleController := scheduleFactory.GetScheduleController()

	controllers := common.Controllers{
		CircuitController:    *circuitController,
		DepartmentController: *departmentController,
		ScheduleController:   *scheduleController,
	}

	routes.PublicRoutes(app, &controllers)

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

func getRabbitWorker() workers.Worker {
	rabbitWorker, err := rabbit.NewRabbitConnection()

	if err != nil {
		panic(err)
	}

	return rabbitWorker
}

func getHelpers(rabbitWorker workers.Worker) helpers.Helpers {
	var loggerCommunication irabbit.RabbitCommunication = rabbit.NewLogRabbitCommunication(rabbitWorker)
	loggerHelper := helpers.NewLogHelper(loggerCommunication)
	helpers := helpers.Helpers{LogHelper: *loggerHelper}

	return helpers
}
