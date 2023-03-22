package helpers

import (
	"monitoring-service/models"
	"monitoring-service/rabbit"
	irabbit "monitoring-service/rabbit/interfaces"
)

type LogHelper struct {
	rabbitCommunication irabbit.RabbitCommunication
}

func NewLogHelper(rabbitCommunication irabbit.RabbitCommunication) *LogHelper {
	return &LogHelper{rabbitCommunication: rabbitCommunication}
}

func (logHelper *LogHelper) SendLog(log models.LoggingModel) {
	jsonLog, logErr := rabbit.ConvertLoggingModelToByteSlice(log)

	if logErr != nil {
		panic(logErr)
	}

	logHelper.rabbitCommunication.Send(jsonLog)
}
