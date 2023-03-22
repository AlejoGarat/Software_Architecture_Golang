package read

type CircuitVoteCoverage struct {
	CircuitId  string     `json:"circuit_id" bson:"circuit_id"`
	AgeVote    AgeVote    `json:"age_vote" bson:"age_vote"`
	GenderVote GenderVote `json:"gender_vote" bson:"gender_vote"`
}

type AgeVote struct {
	CircuitVotersPerAge []CircuitVotersPerAge `json:"voters_age" bson:"voters_age"`
	CircuitVotesPerAge  []CircuitVotesPerAge  `json:"votes_age" bson:"votes_age"`
}

type GenderVote struct {
	CircuitVotersPerGender []CircuitVotersPerGender `json:"voters_gender" bson:"voters_gender"`
	CircuitVotesPerGender  []CircuitVotesPerGender  `json:"votes_gender" bson:"votes_gender"`
}

type CircuitVotersPerAge struct {
	Age    int `json:"age" bson:"age"`
	Voters int `json:"voters" bson:"voters"`
}

type CircuitVotersPerGender struct {
	Gender int `json:"gender" bson:"gender"`
	Votes  int `json:"votes" bson:"votes"`
}

type CircuitVotesPerAge struct {
	Age   int `json:"age" bson:"age"`
	Votes int `json:"votes" bson:"votes"`
}

type CircuitVotesPerGender struct {
	Gender int `json:"gender" bson:"gender"`
	Votes  int `json:"votes" bson:"votes"`
}

type Circuit struct {
	CircuitId string `json:"circuit_id" bson:"circuit_id"`
}
