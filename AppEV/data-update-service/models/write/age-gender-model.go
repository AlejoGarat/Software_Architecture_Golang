package write

type AgeVote struct {
	CircuitVotesPerAge []VotesPerAge `json:"votes_age" bson:"votes_age"`
}

type GenderVote struct {
	CircuitVotesPerGender []VotesPerGender `json:"votes_gender" bson:"votes_gender"`
}

type VotesPerAge struct {
	Age   int `json:"age" bson:"age"`
	Votes int `json:"votes" bson:"votes"`
}

type VotesPerGender struct {
	Gender string `json:"gender" bson:"gender"`
	Votes  int    `json:"votes" bson:"votes"`
}
