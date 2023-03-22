package helpers

import (
	"votation-service/models/write"
	"votation-service/rabbit"
	irabbit "votation-service/rabbit/interfaces"
)

type MessageHelper struct {
	rabbitCommunication irabbit.RabbitCommunication
}

func NewMessageHelper(rabbitCommunication irabbit.RabbitCommunication) *MessageHelper {
	return &MessageHelper{rabbitCommunication: rabbitCommunication}
}

func (messageHelper *MessageHelper) SendMessage(constancy write.Constancy) error {
	jsonMessage, messageErr := rabbit.ConvertConstancyToByteSlice(constancy)

	if messageErr != nil {
		panic(messageErr)
	}

	return messageHelper.rabbitCommunication.Send(jsonMessage)
}
