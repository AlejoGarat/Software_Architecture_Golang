package rabbit

import (
	"consultant-service/models/write"
	workers "consultant-service/rabbit/workers"
	"encoding/json"
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

func ConvertModelToByteSlice(loggingModel write.LoggingModel) ([]byte, error) {
	jsonLog, err := json.Marshal(loggingModel)

	if err != nil {
		return []byte{}, err
	}

	return jsonLog, nil
}
