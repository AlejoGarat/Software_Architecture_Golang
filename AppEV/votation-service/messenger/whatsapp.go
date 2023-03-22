package messenger

import (
	"votation-service/models/write"
	workers "votation-service/rabbit/workers"
)

type Whatsapp struct{}

func NewWhatsapp() *Whatsapp {
	return &Whatsapp{}
}

func (whatsapp *Whatsapp) SendMessage(constancy write.Constancy, worker workers.Worker) error {
	return nil
}
