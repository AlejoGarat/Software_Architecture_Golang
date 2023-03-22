package rabbit

import (
	"encoding/json"
	"votation-service/models/write"
	workers "votation-service/rabbit/workers"
)

type MessageRabbitCommunication struct {
	rabbitWorker workers.Worker
}

func NewMessageRabbitCommunication(rabbitWorker workers.Worker) *MessageRabbitCommunication {
	return &MessageRabbitCommunication{rabbitWorker: rabbitWorker}
}

func (messageRabbitCommunication MessageRabbitCommunication) Send(json []byte) error {
	err := messageRabbitCommunication.rabbitWorker.Send("sms-queue", json)

	if err != nil {
		return err
	}

	return nil
}

func ConvertConstancyToByteSlice(constancy write.Constancy) ([]byte, error) {
	jsonLog, err := json.Marshal(constancy)

	if err != nil {
		return []byte{}, err
	}

	return jsonLog, nil
}
