package filters

import (
	"errors"
	idataaccess "votation-service/dataaccess/interfaces"
	"votation-service/models/write"
)

const (
	unique = "unique"
)

type UniqueVotationMode struct {
	voteRepository idataaccess.VoteRepository
}

func NewUniqueVotationMode(voteRepository idataaccess.VoteRepository) *UniqueVotationMode {
	return &UniqueVotationMode{voteRepository: voteRepository}
}

func (uniqueVotationMode *UniqueVotationMode) Filter(vote write.Vote) error {
	election, err := uniqueVotationMode.voteRepository.GetElectionById(vote.ElectionId)

	if err != nil {
		return err
	}

	voterHasVoted := uniqueVotationMode.voteRepository.ExistsVoteOfVoter(vote.VoterDocument, vote.ElectionId)

	if election.VotationMode == unique && voterHasVoted {
		return errors.New("You cant vote more than one time in this election.")
	}

	return nil
}
