package write

type FrequentVotationSchedules struct {
	ElectionId string     `json:"election_id" bson:"election_id"`
	Schedules  []HourVote `json:"info" bson:"info"`
}

type HourVote struct {
	Hour  string `json:"hour" bson:"hour"`
	Votes int    `json:"votes" bson:"votes"`
}
