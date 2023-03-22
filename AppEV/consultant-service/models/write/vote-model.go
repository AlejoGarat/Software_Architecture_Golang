package write

import (
	"time"
)

type ElectionId = string
type CircuitId = string
type CandidateId = string
type VoterId = string

type Vote struct {
	CreatedAt   time.Time   `json:"created_at"`
	ElectionId  ElectionId  `json:"election_id"`
	CircuitId   CircuitId   `json:"circuit_id"`
	CandidateId CandidateId `json:"candidate_id"`
	VoterId     VoterId     `json:"voter_id"`
}
