package senders

import (
	"encoding/json"
	"message-sender/models"
	"message-sender/workers"
)

func SendLog(log models.LoggingModel, worker workers.Worker) error {
	jsonLog, err := convertLoggingModelToByteSlice(log)

	if err != nil {
		return err
	}

	err = worker.Send("log_queue", jsonLog)

	if err != nil {
		return err
	}

	return nil
}

func convertLoggingModelToByteSlice(loggingModel models.LoggingModel) ([]byte, error) {
	jsonLog, err := json.Marshal(loggingModel)

	if err != nil {
		return []byte{}, err
	}

	return jsonLog, nil
}
