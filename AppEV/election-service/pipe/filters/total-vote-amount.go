package filters

import (
	idataaccess "election-service/dataaccess/interfaces"
	"election-service/models/write"
	"errors"
)

type TotalVoteAmount struct {
	repository idataaccess.ElectionRepository
}

func NewTotalVoteAmount(repository idataaccess.ElectionRepository) *TotalVoteAmount {
	return &TotalVoteAmount{repository: repository}
}

func (ced TotalVoteAmount) Filter(completeElection write.CompleteElection) error {
	votes, err := ced.repository.GetVotes(completeElection.Id)

	if err != nil {
		return err
	}

	if len(votes) > len(completeElection.Voters) {
		return errors.New("There can't be more votes than enables voters.")
	}

	return nil
}
