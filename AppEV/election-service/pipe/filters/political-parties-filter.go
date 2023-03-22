package filters

import (
	"election-service/models/write"
	"errors"
)

type PoliticalPatrties struct {
}

func NewPoliticalParties() *PoliticalPatrties {
	return &PoliticalPatrties{}
}

func (pp PoliticalPatrties) Filter(completeElection write.CompleteElection) error {
	if completeElection.PoliticalParties != nil || len(completeElection.PoliticalParties) > 0 {
		return nil
	}

	return errors.New("There must be at least one political party.")
}
