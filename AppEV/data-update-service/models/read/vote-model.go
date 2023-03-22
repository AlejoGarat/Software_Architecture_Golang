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
	ElectionId    ElectionId `json:"election_id"`
	VoterDocument VoterId    `json:"voter_id"`
	CircuitId     CircuitId  `json:"circuit_id"`
	VoteId        VoteId     `json:"id"`
	Info          Info       `json:"info_struct" bson:"info_struct"`
}

type VoteGet struct {
	VoterDocument VoterId     `json:"voter_id" bson:"voter_id"`
	CircuitId     CircuitId   `json:"circuit_id" bson:"circuit_id"`
	CandidateId   CandidateId `json:"candidate_id" bson:"candidate_id"`
}
