package filters

import (
	"errors"
	idataaccess "votation-service/dataaccess/interfaces"
	"votation-service/models/write"
)

type EnabledVoter struct {
	voteRepository idataaccess.VoteRepository
}

func NewEnabledVoter(voteRepository idataaccess.VoteRepository) *EnabledVoter {
	return &EnabledVoter{voteRepository: voteRepository}
}

func (enabledVoter *EnabledVoter) Filter(vote write.Vote) error {
	_, err := enabledVoter.voteRepository.GetVoterById(vote.VoterDocument, vote.ElectionId)

	if err != nil {
		return errors.New("The voters document does not exists.")
	}

	return nil
}
