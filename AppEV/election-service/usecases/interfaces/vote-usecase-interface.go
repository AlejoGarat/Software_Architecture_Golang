package interfaces

import "election-service/models/read"

type VoteUseCase interface {
	GetVoterVotingSchedules(string, string) ([]read.Info, error)
}
