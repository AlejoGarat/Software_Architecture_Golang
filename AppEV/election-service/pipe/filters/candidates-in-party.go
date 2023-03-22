package filters

import (
	"election-service/models/write"
	"errors"
)

type CandidatesInParty struct {
}

func NewCandidatesInParty() *CandidatesInParty {
	return &CandidatesInParty{}
}

func (cip CandidatesInParty) Filter(completeElection write.CompleteElection) error {

	for _, candidate := range completeElection.Candidates {
		partyAmount := 0
		for _, party := range completeElection.PoliticalParties {
			if party.Id == candidate.PoliticalPartyId {
				partyAmount = partyAmount + 1

				if partyAmount > 1 {
					return errors.New("Every candidate must be in exactly one political party.")
				}
			}
		}
	}

	return nil
}
