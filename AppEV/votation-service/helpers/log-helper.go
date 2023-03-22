package helpers

import (
	"votation-service/models/write"
	"votation-service/rabbit"
	irabbit "votation-service/rabbit/interfaces"
)

type LogHelper struct {
	rabbitCommunication irabbit.RabbitCommunication
}

func NewLogHelper(rabbitCommunication irabbit.RabbitCommunication) *LogHelper {
	return &LogHelper{rabbitCommunication: rabbitCommunication}
}

func (logHelper *LogHelper) SendLog(log write.LoggingModel) error {
	jsonLog, logErr := rabbit.ConvertModelToByteSlice(log)

	if logErr != nil {
		panic(logErr)
	}

	return logHelper.rabbitCommunication.Send(jsonLog)
}
