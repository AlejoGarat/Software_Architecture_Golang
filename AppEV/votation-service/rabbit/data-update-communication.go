package rabbit

import (
	"encoding/json"
	"votation-service/models/write"
	workers "votation-service/rabbit/workers"
)

type DataUpdateRabbitCommunication struct {
	rabbitWorker workers.Worker
}

func NewDataUpdateRabbitCommunication(rabbitWorker workers.Worker) *DataUpdateRabbitCommunication {
	return &DataUpdateRabbitCommunication{rabbitWorker: rabbitWorker}
}

func (dataUpdateRabbitCommunication DataUpdateRabbitCommunication) Send(json []byte) error {
	err := dataUpdateRabbitCommunication.rabbitWorker.Send("vote-update-queue", json)

	if err != nil {
		return err
	}

	return nil
}

func ConvertVoteToByteSlice(vote write.Vote) ([]byte, error) {
	jsonDataUpdate, err := json.Marshal(vote)

	if err != nil {
		return []byte{}, err
	}

	return jsonDataUpdate, nil
}
