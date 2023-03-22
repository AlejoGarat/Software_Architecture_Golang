package filters

import (
	"errors"
	idataaccess "votation-service/dataaccess/interfaces"
	"votation-service/models/write"
)

type CircuitVoter struct {
	voteRepository idataaccess.VoteRepository
}

func NewCircuitVoter(voteRepository idataaccess.VoteRepository) *CircuitVoter {
	return &CircuitVoter{voteRepository: voteRepository}
}

func (circuitVoter *CircuitVoter) Filter(vote write.Vote) error {
	voter, err := circuitVoter.voteRepository.GetVoterById(vote.VoterDocument, vote.ElectionId)

	if err != nil {
		return errors.New("he voters document does not exists.")
	}

	if voter.Circuit != vote.CircuitId {
		return errors.New("The circuit sent is not the correct voters circuit.")
	}
	return nil
}
