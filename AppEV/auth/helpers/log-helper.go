package helpers

import (
	"auth/models/write"
	"auth/rabbit"
	irabbit "auth/rabbit/interfaces"
	"errors"
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
		return errors.New("error while converting logging model to byte slice")
	}

	return logHelper.rabbitCommunication.Send(jsonLog)
}
