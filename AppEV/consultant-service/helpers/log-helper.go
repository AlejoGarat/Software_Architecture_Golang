package helpers

import (
	"consultant-service/models/write"
	"consultant-service/rabbit"
	irabbit "consultant-service/rabbit/interfaces"
)

type LogHelper struct {
	rabbitCommunication irabbit.RabbitCommunication
}

func NewLogHelper(rabbitCommunication irabbit.RabbitCommunication) *LogHelper {
	return &LogHelper{rabbitCommunication: rabbitCommunication}
}

func (logHelper *LogHelper) SendLog(log write.LoggingModel) {
	jsonLog, logErr := rabbit.ConvertModelToByteSlice(log)

	if logErr != nil {
		panic(logErr)
	}

	logHelper.rabbitCommunication.Send(jsonLog)
}
