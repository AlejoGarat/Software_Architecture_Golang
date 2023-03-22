package rabbit

import (
	"encoding/json"
	"monitoring-service/models"
	workers "monitoring-service/rabbit/workers"
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

func ConvertLoggingModelToByteSlice(loggingModel models.LoggingModel) ([]byte, error) {
	jsonLog, err := json.Marshal(loggingModel)

	if err != nil {
		return []byte{}, err
	}

	return jsonLog, nil
}
