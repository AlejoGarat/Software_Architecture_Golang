package write

import "time"

type ElectionId = string
type CircuitId = string
type VoterId = string
type CandidateId = string
type VoteId = string
type Time = time.Time
type ProcessingTime = string

type Info struct {
	Date string `bson:"date"`
	Hour string `bson:"hour"`
}

type Vote struct {
	ElectionId     ElectionId     `json:"election_id" bson:"election_id"`
	VoterDocument  VoterId        `json:"voter_id" bson:"voter_id"`
	CircuitId      CircuitId      `json:"circuit_id" bson:"circuit_id"`
	CandidateId    CandidateId    `json:"candidate_id" bson:"candidate_id"`
	VoteId         VoteId         `json:"id" bson:"id"`
	Info           Info           `json:"info_struct" bson:"info_struct"`
	InfoArr        []Info         `bson:"info"`
	StartingTime   Time           `bson:"starting_time"`
	EndingTime     Time           `bson:"ending_time"`
	ProcessingTime ProcessingTime `bson:"processing_time"`
}
