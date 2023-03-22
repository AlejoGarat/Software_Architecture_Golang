package write

type DepartmentVoteCoverage struct {
	DepartmentVotesPerAge    []VotesPerAge    `json:"votes_age" bson:"votes_age"`
	DepartmentVotesPerGender []VotesPerGender `json:"votes_gender" bson:"votes_gender"`
}
