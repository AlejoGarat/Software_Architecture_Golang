package interfaces

import (
	"votation-service/models/read"
	"votation-service/models/write"
)

type VoteUseCase interface {
	AddVote(write.Vote)
	SendMailConstancy(read.ConstancyRequest) error
	UpdateConstancyDBData(write.ConstancyDBData, string) error
}
