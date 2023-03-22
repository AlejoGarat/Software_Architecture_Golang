package filters

import (
	"errors"
	"reflect"
	idataaccess "votation-service/dataaccess/interfaces"
	"votation-service/models/write"
)

type UniqueCandidate struct {
	voteRepository idataaccess.VoteRepository
}

func NewUniqueCandidate(voteRepository idataaccess.VoteRepository) *UniqueCandidate {
	return &UniqueCandidate{voteRepository: voteRepository}
}

func (uniqueCandidate *UniqueCandidate) Filter(vote write.Vote) error {
	var strType string

	if reflect.TypeOf(vote.CandidateId) != reflect.TypeOf(strType) {
		return errors.New("The candidate id has an incorrect format.")
	}

	return nil
}
