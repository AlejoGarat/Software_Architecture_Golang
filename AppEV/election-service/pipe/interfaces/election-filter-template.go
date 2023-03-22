package interfaces

import "election-service/models/write"

type ElectionFilter interface {
	Filter(write.CompleteElection) error
}
