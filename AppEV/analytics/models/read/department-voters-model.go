package read

type DepartmentVoteCoverage struct {
	DepartmentName string     `json:"circuit_id" bson:"circuit_id"`
	AgeVote        AgeVote    `json:"age_vote" bson:"age_vote"`
	GenderVote     GenderVote `json:"gender_vote" bson:"gender_vote"`
}

type DepartmentAgeVote struct {
	CircuitVotersPerAge []CircuitVotersPerAge `json:"voters_age" bson:"voters_age"`
	CircuitVotesPerAge  []CircuitVotesPerAge  `json:"votes_age" bson:"votes_age"`
}

type DepartmentGenderVote struct {
	CircuitVotersPerGender []CircuitVotersPerGender `json:"voters_gender" bson:"voters_gender"`
	CircuitVotesPerGender  []CircuitVotesPerGender  `json:"votes_gender" bson:"votes_gender"`
}

type DepartmentVotersPerAge struct {
	Age   int `json:"name" bson:"name"`
	Votes int `json:"votes" bson:"votes"`
}

type DepartmentVotersPerGender struct {
	Gender int `json:"gender" bson:"gender"`
	Votes  int `json:"votes" bson:"votes"`
}

type DepartmentVotesPerAge struct {
	Age   int `json:"name" bson:"name"`
	Votes int `json:"votes" bson:"votes"`
}

type DepartmentVotesPerGender struct {
	Gender int `json:"gender" bson:"gender"`
	Votes  int `json:"votes" bson:"votes"`
}
