package rabbit

import (
	workers "election-service/rabbit/workers"
)

type LogRabbitCommunication struct {
	rabbitWorker workers.Worker
}

func NewLogRabbitCommunication(rabbitWorker workers.Worker) *LogRabbitCommunication {
	return &LogRabbitCommunication{rabbitWorker: rabbitWorker}
}

func (logRabbitCommunication LogRabbitCommunication) Send(json []byte) error {
	err := logRabbitCommunication.rabbitWorker.Send("log-queue", json)

	if err != nil {
		return err
	}

	return nil
}

func (logRabbitCommunication LogRabbitCommunication) SendSignal(json []byte) error {
	err := logRabbitCommunication.rabbitWorker.Send("update-data-queue", json)

	if err != nil {
		return err
	}

	return nil
}
