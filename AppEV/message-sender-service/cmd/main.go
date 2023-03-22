package main

import (
	"encoding/json"
	"log"
	"message-sender/models"
	"message-sender/senders"
	"message-sender/workers"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	openENV()

	rabbitWorker := getRabbitWorker()

	listenForSMS(rabbitWorker)

	rabbitWorker.Close()
}

func listenForSMS(worker workers.Worker) {

	for {
		dequeueConstancyMessage(worker)
	}
}

func dequeueConstancyMessage(worker workers.Worker) {
	var log models.LoggingModel

	worker.Listen(100, "sms-queue", func(message []byte) error {
		var constancy models.Constancy

		err := json.Unmarshal(message, &constancy)

		if err != nil {
			log = models.LoggingModel{Type: "Error", Operation: "Dequeue of Message Constancy Election", Actor: "N/A", Description: err.Error()}
			senders.SendLog(log, worker)
		}

		senders.SendSms(constancy)
		return nil
	})
}

func openENV() {
	dotenvErr := godotenv.Load("../.env")

	if dotenvErr != nil {
		log.Fatalf("Error loading .env file")
	}
}

func getRabbitWorker() workers.Worker {
	rabbit := os.Getenv("RABBIT")
	rabbitWorker, err := workers.BuildRabbitWorker(rabbit)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return rabbitWorker
}
