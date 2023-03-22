package interfaces

import "election-service/models/read"

type VoteRepository interface {
	GetVoterVotingSchedules(string, string) ([]read.Info, error)
}
