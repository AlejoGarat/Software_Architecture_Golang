package rabbit

import (
	"encoding/json"
	"votation-service/models/write"
	workers "votation-service/rabbit/workers"
)

type MailRabbitCommunication struct {
	rabbitWorker workers.Worker
}

func NewMailRabbitCommunication(rabbitWorker workers.Worker) *MailRabbitCommunication {
	return &MailRabbitCommunication{rabbitWorker: rabbitWorker}
}

func (mailRabbitCommunication MailRabbitCommunication) Send(json []byte) error {
	err := mailRabbitCommunication.rabbitWorker.Send("mail-constancy-queue", json)

	if err != nil {
		return err
	}

	return nil
}

func ConvertMailConstancyToByteSlice(constancy write.Constancy) ([]byte, error) {
	jsonLog, err := json.Marshal(constancy)

	if err != nil {
		return []byte{}, err
	}

	return jsonLog, nil
}
