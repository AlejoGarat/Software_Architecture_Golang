package filters

import (
	"election-service/models/write"
	"errors"
	"strings"
)

type CorrectVotationMode struct {
}

func NewCorrectVotationMode() *CorrectVotationMode {
	return &CorrectVotationMode{}
}

const (
	multiple = "multiple"
	unique   = "unique"
)

func (cip CorrectVotationMode) Filter(completeElection write.CompleteElection) error {

	votationMode := strings.ToLower(completeElection.VotationMode)

	if votationMode != multiple && votationMode != unique {
		return errors.New("Invalid votation mode. Must be 'Multiple' or 'Unique'.")
	}

	return nil
}
