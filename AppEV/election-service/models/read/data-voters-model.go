package read

type CircuitVoteCoverage struct {
	CircuitId     string `json:"circuit_id" bson:"circuit_id"`
	CircuitAge    CircuitAge
	CircuitGender CircuitGender
}

type CircuitAge struct {
	AgeVoters []VotersAge   `json:"voters_age" bson:"voters_age"`
	AgeVotes  []VotesPerAge `json:"votes_age" bson:"votes_age"`
}

type CircuitGender struct {
	GenderVoters []VotersGender   `json:"voters_gender" bson:"voters_gender"`
	GenderVotes  []VotesPerGender `json:"votes_gender" bson:"votes_gender"`
}

type DepartmentVoteCoverage struct {
	DepartmentName string           `json:"department" bson:"department"`
	AgeVoters      []VotersAge      `json:"voters_age" bson:"voters_age"`
	GenderVoters   []VotersGender   `json:"voters_gender" bson:"voters_gender"`
	AgeVotes       []VotesPerAge    `json:"votes_age" bson:"votes_age"`
	GenderVotes    []VotesPerGender `json:"votes_gender" bson:"votes_gender"`
}

type VotersAge struct {
	Age    int `json:"age" bson:"age"`
	Voters int `json:"voters" bson:"voters"`
}

type VotersGender struct {
	Gender string `json:"gender" bson:"gender"`
	Votes  int    `json:"voters" bson:"voters"`
}

type VotesPerAge struct {
	Age   int `json:"age" bson:"age"`
	Votes int `json:"votes" bson:"votes"`
}

type VotesPerGender struct {
	Gender string `json:"gender" bson:"gender"`
	Votes  int    `json:"votes" bson:"votes"`
}

type Circuit struct {
	CircuitId string `json:"id" bson:"id"`
}
