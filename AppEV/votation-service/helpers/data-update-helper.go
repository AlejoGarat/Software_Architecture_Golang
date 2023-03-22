package helpers

import (
	"votation-service/models/write"
	"votation-service/rabbit"
	irabbit "votation-service/rabbit/interfaces"
)

type DataUpdateHelper struct {
	rabbitCommunication irabbit.RabbitCommunication
}

func NewDataUpdateHelper(rabbitCommunication irabbit.RabbitCommunication) *DataUpdateHelper {
	return &DataUpdateHelper{rabbitCommunication: rabbitCommunication}
}

func (dataUpdateHelper *DataUpdateHelper) SendVote(vote write.Vote) error {
	jsonVote, jsonErr := rabbit.ConvertVoteToByteSlice(vote)

	if jsonErr != nil {
		panic(jsonErr)
	}

	return dataUpdateHelper.rabbitCommunication.Send(jsonVote)
}
