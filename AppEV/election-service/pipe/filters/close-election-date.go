package filters

import (
	"election-service/models/write"
	"errors"
	"time"
)

type CloseElectionDate struct {
}

func NewCloseElectionDate() *CloseElectionDate {
	return &CloseElectionDate{}
}

func (ced CloseElectionDate) Filter(completeElection write.CompleteElection) error {
	if completeElection.EndDate.Before(time.Now().Add(-time.Hour * 1)) {
		return errors.New("Election cannot end before votation period ended.")
	}

	return nil
}
