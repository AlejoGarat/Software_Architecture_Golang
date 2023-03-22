package usecases

import (
	idataaccess "election-service/dataaccess/interfaces"
	"election-service/helpers"
	"election-service/models/read"
	"election-service/models/write"
)

type VoteUseCase struct {
	voteRepository idataaccess.VoteRepository
	helpers        helpers.Helpers
}

func NewVoteUseCase(voteRepository idataaccess.VoteRepository, helpers helpers.Helpers) *VoteUseCase {
	return &VoteUseCase{voteRepository: voteRepository, helpers: helpers}
}

func (voteUseCase *VoteUseCase) GetVoterVotingSchedules(electionId string, voterId string) ([]read.Info, error) {
	var log write.LoggingModel
	voterSchedules, err := voteUseCase.voteRepository.GetVoterVotingSchedules(electionId, voterId)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Get Voter Voting Schedules", Actor: "Electoral Authority", Description: err.Error()}
		voteUseCase.helpers.LogHelper.SendLog(log)
	}

	return voterSchedules, err
}
