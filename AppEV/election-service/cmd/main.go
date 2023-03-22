package main

import (
	"context"
	"election-service/api/common"
	"election-service/api/configs"
	"election-service/api/middlewares"
	"election-service/api/routes"
	"election-service/datasources"
	"election-service/factory"
	"election-service/helpers"
	"election-service/models/write"
	"election-service/rabbit"
	irabbit "election-service/rabbit/interfaces"
	"election-service/rabbit/workers"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	mongoConnection = "mongodb://localhost:27017/"
	electionsDB     = "ElectoralDB"
	redisConnection = "redis://localhost:6379/"
)

func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	middlewares.FiberMiddleware(app)

	openENV()

	rabbitWorker := getRabbitWorker()

	helpers := getHelpers(rabbitWorker)

	var log write.LoggingModel

	electoral_db := os.Getenv("ELECTORAL_DATABASE")

	election_db := os.Getenv("ELECTION_DATABASE")

	mongoAddress := os.Getenv("MONGO_CLIENT")

	redisAddress := os.Getenv("REDIS_CLIENT")

	mongoClient, err := getMongoClient(mongoAddress)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Mongo Connection", Actor: "Electoral Authority", Description: err.Error()}
		helpers.LogHelper.SendLog(log)
	}

	defer mongoClient.Disconnect(context.TODO())

	redisClient, err := getRedisClient(redisAddress)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Redis Connection", Actor: "Electoral Authority", Description: err.Error()}
		helpers.LogHelper.SendLog(log)
	}

	defer redisClient.Close()

	electionFactory := factory.NewElectionFactory(mongoClient, redisClient, electoral_db, helpers)
	alertFactory := factory.NewAlertFactory(redisClient, *helpers)
	circuitFactory := factory.NewCircuitFactory(mongoClient, electoral_db, helpers)
	departmentFactory := factory.NewDepartmentFactory(mongoClient, electoral_db, helpers)
	scheduleFactory := factory.NewScheduleFactory(mongoClient, electoral_db, helpers)
	voteFactory := factory.NewVoteFactory(mongoClient, election_db, helpers)

	electionController := electionFactory.GetElectionController()
	alertController := alertFactory.GetAlertController()
	circuitController := circuitFactory.GetCircuitController()
	departmentController := departmentFactory.GetDepartmentController()
	scheduleController := scheduleFactory.GetScheduleController()
	voteController := voteFactory.GetVoteController()

	controllers := common.Controllers{
		ElectionController:   *electionController,
		CircuitController:    *circuitController,
		DepartmentController: *departmentController,
		ScheduleController:   *scheduleController,
		AlertController:      *alertController,
		VoteController:       *voteController,
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
