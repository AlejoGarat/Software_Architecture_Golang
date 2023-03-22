package helpers

import (
	"analytics/models/write"
	irabbit "analytics/rabbit/interfaces"
	"encoding/json"
	"log"
)

type LogHelper struct {
	rabbitCommunication irabbit.RabbitCommunication
}

func NewLogHelper(rabbitCommunication irabbit.RabbitCommunication) *LogHelper {
	return &LogHelper{rabbitCommunication: rabbitCommunication}
}

func (logHelper LogHelper) SendLog(logg write.LoggingModel) {
	jsonLog, logErr := convertLoggingModelToByteSlice(logg)

	if logErr != nil {
		log.Fatalf(logErr.Error())
	}

	err := logHelper.rabbitCommunication.Send(jsonLog)

	if err != nil {
		log.Fatalf(err.Error())
	}
}

func convertLoggingModelToByteSlice(loggingModel write.LoggingModel) ([]byte, error) {
	jsonLog, err := json.Marshal(loggingModel)

	if err != nil {
		return []byte{}, err
	}

	return jsonLog, nil
}
