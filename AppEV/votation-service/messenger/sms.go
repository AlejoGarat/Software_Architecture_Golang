package messenger

import (
	"votation-service/helpers"
	"votation-service/models/write"
)

type Sms struct {
	messageHelper helpers.MessageHelper
}

func NewSms(messageHelper helpers.MessageHelper) *Sms {
	return &Sms{messageHelper: messageHelper}
}

func (sms *Sms) SendMessage(constancy write.Constancy) error {
	return sms.messageHelper.SendMessage(constancy)
}
