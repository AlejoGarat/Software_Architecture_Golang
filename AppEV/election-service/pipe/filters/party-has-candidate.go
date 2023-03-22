package filters

import (
	"election-service/models/write"
	"errors"
)

type PartyHasCanidate struct {
}

func NewPartyHasCandidate() *PartyHasCanidate {
	return &PartyHasCanidate{}
}

func (phc PartyHasCanidate) Filter(completeElection write.CompleteElection) error {
	for _, party := range completeElection.PoliticalParties {
		hasCandidate := false
		for _, candidate := range completeElection.Candidates {
			if party.Id == candidate.PoliticalPartyId {
				hasCandidate = true
				break
			}
		}

		if !hasCandidate {
			return errors.New("Every political party must have at least one candidate.")
		}
	}

	return nil
}
