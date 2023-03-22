package interfaces

import "data-update-service/models/read"

type VoteDataUseCase interface {
	UpdateVoteData(read.Vote) error
	StartUpdateCrone(read.Election) error
}
