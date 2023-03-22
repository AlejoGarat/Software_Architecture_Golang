package main

import (
	"communication-sender/models"
	"communication-sender/senders"
	"communication-sender/workers"
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	openENV()

	rabbitWorker := getRabbitWorker()

	go listenForMailConstancy(rabbitWorker)

	go listenForOpenCertificate(rabbitWorker)

	go listenForCloseCertificate(rabbitWorker)

	listenForConfigAlerts(rabbitWorker)

	rabbitWorker.Close()
}

func listenForMailConstancy(worker workers.Worker) {

	for {
		dequeueMailConstancy(worker)
	}
}

func listenForOpenCertificate(worker workers.Worker) {

	for {
		dequeueOpenCertificate(worker)
	}
}

func listenForCloseCertificate(worker workers.Worker) {

	for {
		dequeueCloseCertificate(worker)
	}
}

func listenForConfigAlerts(worker workers.Worker) {

	for {
		dequeueMailAlertConfig(worker)
	}
}

func dequeueMailConstancy(worker workers.Worker) {
	var log models.LoggingModel

	worker.Listen(50, "mail-constancy-queue", func(message []byte) error {
		var constancy models.Constancy

		err := json.Unmarshal(message, &constancy)

		if err != nil {
			log = models.LoggingModel{Type: "Error", Operation: "Dequeue of Mail Constancy Election", Actor: "N/A", Description: err.Error()}
			senders.SendLog(log, worker)
		}

		senders.SendMail(constancy)
		return nil
	})
}

func dequeueOpenCertificate(worker workers.Worker) {
	var log models.LoggingModel

	worker.Listen(50, "open-certificate-queue", func(message []byte) error {
		var openCertificate models.StartCertificate

		err := json.Unmarshal(message, &openCertificate)

		if err != nil {
			log = models.LoggingModel{Type: "Error", Operation: "Dequeue of Open Certificate Election", Actor: "N/A", Description: err.Error()}
			senders.SendLog(log, worker)
		}

		senders.SendOpenCertificate(openCertificate)
		return nil
	})
}

func dequeueCloseCertificate(worker workers.Worker) {
	var log models.LoggingModel

	worker.Listen(50, "close-certificate-queue", func(message []byte) error {
		var closeCertificate models.CloseCertificate

		err := json.Unmarshal(message, &closeCertificate)

		if err != nil {
			log = models.LoggingModel{Type: "Error", Operation: "Dequeue of Close Certificate Election", Actor: "N/A", Description: err.Error()}
			senders.SendLog(log, worker)
		}

		senders.SendCloseCertificate(closeCertificate)
		return nil
	})
}

func dequeueMailAlertConfig(worker workers.Worker) {
	worker.Listen(50, "mail-config-queue", func(message []byte) error {
		senders.SendMailAlert(string(message))
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
