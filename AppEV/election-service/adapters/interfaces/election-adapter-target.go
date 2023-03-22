package interfaces

import "election-service/models/write"

type ElectionAdapterTarget interface {
	ConvertElection([]byte) (write.CompleteElection, error)
}
