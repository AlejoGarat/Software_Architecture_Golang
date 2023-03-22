package main

import (
	"fmt"
	"log"
	"log-service/logger"
	ilogger "log-service/logger/interfaces"
	"log-service/workers"
	"os"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	worker, err := workers.BuildRabbitWorker("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println(err.Error())
	}

	var logger ilogger.Logger = logger.NewLogger()

	listenForLogs(worker, logger)

	worker.Close()
}

func listenForLogs(worker workers.Worker, logger ilogger.Logger) {
	for {
		worker.Listen(50, "log-queue", func(message []byte) error {
			logger.Log(string(message))
			return nil
		})
	}
}
