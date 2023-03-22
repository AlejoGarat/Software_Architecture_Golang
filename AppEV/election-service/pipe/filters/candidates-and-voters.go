package filters

import (
	"election-service/models/write"
	"errors"
)

type CandidatesAndVoters struct {
}

func NewCandidatesAndVoters() *CandidatesAndVoters {
	return &CandidatesAndVoters{}
}

func (cav CandidatesAndVoters) Filter(completeElection write.CompleteElection) error {
	if completeElection.Candidates == nil || len(completeElection.Candidates) == 0 {
		return errors.New("There must be at least one candidate.")
	}

	if completeElection.Voters == nil || len(completeElection.Voters) == 0 {
		return errors.New("There must be at least one voter.")
	}

	return nil
}
