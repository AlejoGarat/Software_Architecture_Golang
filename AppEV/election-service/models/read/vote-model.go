package read

type ElectionId = string
type CircuitId = string
type VoterId = string
type CandidateId = string
type VoteId = string

type Info struct {
	Date string `json:"date" bson:"date"`
	Hour string `json:"hour" bson:"hour"`
}

type Vote struct {
	ElectionId    ElectionId  `json:"election_id" bson:"election_id"`
	VoterDocument VoterId     `json:"voter_id" bson:"voter_id"`
	CircuitId     CircuitId   `json:"circuit_id" bson:"circuit_id"`
	CandidateId   CandidateId `json:"candidate_id" bson:"candidate_id"`
	VoteId        VoteId      `json:"id" bson:"id"`
	Info          []Info      `json:"info" bson:"info"`
}
