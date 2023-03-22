package helpers

import (
	"votation-service/models/write"
	"votation-service/rabbit"
	irabbit "votation-service/rabbit/interfaces"
)

type MailHelper struct {
	rabbitCommunication irabbit.RabbitCommunication
}

func NewMailHelper(rabbitCommunication irabbit.RabbitCommunication) *MailHelper {
	return &MailHelper{rabbitCommunication: rabbitCommunication}
}

func (mailHelper *MailHelper) SendMail(constancy write.Constancy) error {
	jsonConstancy, logErr := rabbit.ConvertMailConstancyToByteSlice(constancy)

	if logErr != nil {
		return logErr
	}

	return mailHelper.rabbitCommunication.Send(jsonConstancy)
}
