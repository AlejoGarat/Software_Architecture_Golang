package helpers

import (
	"election-service/models/write"
	irabbit "election-service/rabbit/interfaces"
	"encoding/json"
)

type LogHelper struct {
	rabbitCommunication irabbit.RabbitCommunication
}

func NewLogHelper(rabbitCommunication irabbit.RabbitCommunication) *LogHelper {
	return &LogHelper{rabbitCommunication: rabbitCommunication}
}

func (logHelper *LogHelper) SendLog(log write.LoggingModel) {
	jsonLog, logErr := convertModelToByteSlice(log)

	if logErr != nil {
		panic(logErr)
	}

	logHelper.rabbitCommunication.Send(jsonLog)
}

func (logHelper *LogHelper) SendSignal(jsonElection []byte) {
	logHelper.rabbitCommunication.SendSignal(jsonElection)
}

func convertModelToByteSlice(loggingModel write.LoggingModel) ([]byte, error) {
	jsonLog, err := json.Marshal(loggingModel)

	if err != nil {
		return []byte{}, err
	}

	return jsonLog, nil
}
