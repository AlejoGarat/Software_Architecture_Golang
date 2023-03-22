package filters

import (
	"errors"
	idataaccess "votation-service/dataaccess/interfaces"
	"votation-service/models/write"
)

type CandidateVoter struct {
	voteRepository idataaccess.VoteRepository
}

func NewCandidateVoter(voteRepository idataaccess.VoteRepository) *CandidateVoter {
	return &CandidateVoter{voteRepository: voteRepository}
}

func (candidateVoter *CandidateVoter) Filter(vote write.Vote) error {
	_, err := candidateVoter.voteRepository.GetCandidateById(vote.CandidateId, vote.ElectionId)

	if err != nil {
		return errors.New("This candidate is not a valid candidate in this election.")
	}

	return nil
}
