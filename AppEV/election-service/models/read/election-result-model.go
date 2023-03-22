package read

type ElectionResult struct {
	ElectionVotingData          ElectionVotingData            `json:"election_voting_data"`
	EligibleVotersPerDepartment []StringIntEligibleDepartment `json:"eligible_voters_department" bson:"eligible_voters_department"`
	VotesPerDepartment          []StringIntDepartment         `json:"votes_per_department" bson:"votes_per_department"`
	VotesPerCandidate           []StringIntCandidate          `json:"votes_per_candidate" bson:"votes_per_candidate"`
	VotesPerPoliticalParty      []StringIntParty              `json:"votes_per_political_party" bson:"votes_per_political_party"`
}

type StringIntDepartment struct {
	Id     string `json:"department"`
	Amount int    `json:"amount"`
}

type StringIntEligibleDepartment struct {
	Id     string `json:"eligible_department"`
	Amount int    `json:"amount"`
}

type StringIntParty struct {
	Id     string `json:"party"`
	Amount int    `json:"amount"`
}

type StringIntCandidate struct {
	Id     string `json:"candidate"`
	Amount int    `json:"amount"`
}

type ElectionVotingData struct {
	EligibleVoters  int `json:"eligible_voters" bson:"eligible_voters"`
	TotalVoteAmount int `json:"total_votes" bson:"total_votes"`
}
