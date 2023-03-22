package filters

import (
	"errors"
	"time"
	idataaccess "votation-service/dataaccess/interfaces"
	"votation-service/models/write"
)

type EndedElection struct {
	voteRepository idataaccess.VoteRepository
}

func NewEndedElection(voteRepository idataaccess.VoteRepository) *EndedElection {
	return &EndedElection{voteRepository: voteRepository}
}

func (endedElection *EndedElection) Filter(vote write.Vote) error {
	electionData, err := endedElection.voteRepository.GetElectionById(vote.ElectionId)

	if err != nil {
		return err
	}

	if electionData.End.Before(time.Now()) {
		return errors.New("Election has already ended.")
	}

	return nil
}
