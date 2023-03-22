package rabbit

import (
	"encoding/json"
	workers "monitoring-service/rabbit/workers"
)

type MailRabbitCommunication struct {
	rabbitWorker workers.Worker
}

func NewMailRabbitCommunication(rabbitWorker workers.Worker) *MailRabbitCommunication {
	return &MailRabbitCommunication{rabbitWorker: rabbitWorker}
}

func (mailRabbitCommunication MailRabbitCommunication) Send(json []byte) error {
	err := mailRabbitCommunication.rabbitWorker.Send("mail-config-queue", json)

	if err != nil {
		return err
	}

	return nil
}

func ConvertMessageToByteSlice(message string) ([]byte, error) {
	jsonLog, err := json.Marshal(message)

	if err != nil {
		return []byte{}, err
	}

	return jsonLog, nil
}
