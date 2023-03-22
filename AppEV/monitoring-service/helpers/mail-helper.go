package helpers

import (
	"monitoring-service/rabbit"
	irabbit "monitoring-service/rabbit/interfaces"
)

type MailHelper struct {
	rabbitCommunication irabbit.RabbitCommunication
}

func NewMailHelper(rabbitCommunication irabbit.RabbitCommunication) *MailHelper {
	return &MailHelper{rabbitCommunication: rabbitCommunication}
}

func (mailHelper *MailHelper) SendMail(message string) {
	jsonMail, mailErr := rabbit.ConvertMessageToByteSlice(message)

	if mailErr != nil {
		panic(mailErr)
	}

	mailHelper.rabbitCommunication.Send(jsonMail)
}
