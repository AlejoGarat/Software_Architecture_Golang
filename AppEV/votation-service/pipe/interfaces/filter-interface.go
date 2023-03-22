package interfaces

import "votation-service/models/write"

type VoteFilter interface {
	Filter(write.Vote) error
}
