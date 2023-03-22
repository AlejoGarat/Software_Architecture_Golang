package filters

import (
	"election-service/models/write"
	"errors"
	"time"
)

type ValidDate struct {
}

func NewValidDate() *ValidDate {
	return &ValidDate{}
}

func (vd ValidDate) Filter(completeElection write.CompleteElection) error {
	if completeElection.StartDate.Before(time.Now()) {
		return errors.New("The start date must be greater the actual date.")
	}

	return nil
}
