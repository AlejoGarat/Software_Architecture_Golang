package main

import (
	"context"
	"log"

	"consultant-service/api/configs"
	"consultant-service/api/middlewares"
	"consultant-service/api/routes"
	"consultant-service/datasources"
	"consultant-service/factory"
	"consultant-service/helpers"
	"consultant-service/models/write"
	"consultant-service/rabbit"
	irabbit "consultant-service/rabbit/interfaces"
	"consultant-service/rabbit/workers"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"

	"os"
)

func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	middlewares.FiberMiddleware(app)

	openENV()

	rabbitWorker := getRabbitWorker()

	helpers := getHelpers(rabbitWorker)

	var log write.LoggingModel

	electoral_db := os.Getenv("ELECTORAL_DB")

	mongoAddress := os.Getenv("MONGO_CLIENT")

	mongoClient, err := getMongoClient(mongoAddress)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Mongo Connection", Actor: "Consultant", Description: err.Error()}
		helpers.LogHelper.SendLog(log)
	}

	defer mongoClient.Disconnect(context.TODO())

	redisClient, err := getRedisClient("localhost:6379")

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Redis Connection", Actor: "Consultant", Description: err.Error()}
		helpers.LogHelper.SendLog(log)
	}

	defer redisClient.Close()

	alertFactory := factory.NewAlertFactory(redisClient, *helpers)
	filterFactory := factory.NewFilterFactory(redisClient, *helpers)
	circuitFactory := factory.NewCircuitFactory(mongoClient, electoral_db, helpers)
	departmentFactory := factory.NewDepartmentFactory(mongoClient, electoral_db, helpers)
	scheduleFactory := factory.NewScheduleFactory(mongoClient, electoral_db, helpers)

	alertController := alertFactory.GetAlertController()
	filterController := filterFactory.GetFilterController()
	circuitController := circuitFactory.GetCircuitController()
	departmentController := departmentFactory.GetDepartmentController()
	scheduleController := scheduleFactory.GetScheduleController()

	routes.PublicRoutes(app, filterController, alertController, circuitController, departmentController, scheduleController)

	port := os.Getenv("PORT")

	app.Listen(port)
}

func openENV() {
	dotenvErr := godotenv.Load("../.env")

	if dotenvErr != nil {
		log.Fatalf("Error loading .env file")
	}
}

func getRabbitWorker() workers.Worker {
	rabbitWorker, err := rabbit.NewRabbitConnection()

	if err != nil {
		log.Println(err.Error())
	}

	return rabbitWorker
}

func getHelpers(rabbitWorker workers.Worker) *helpers.Helpers {
	var loggerCommunication irabbit.RabbitCommunication = rabbit.NewLogRabbitCommunication(rabbitWorker)
	loggerHelper := helpers.NewLogHelper(loggerCommunication)
	helpers := helpers.Helpers{LogHelper: *loggerHelper}

	return &helpers
}

func getMongoClient(mongoAddress string) (*mongo.Client, error) {
	return datasources.NewMongoDataSource(mongoAddress)
}

func getRedisClient(redisAddress string) (*redis.Client, error) {
	return datasources.NewRedisDataSource(redisAddress)
}
