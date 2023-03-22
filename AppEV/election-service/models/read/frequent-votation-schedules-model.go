package read

type FrequentVotationSchedules struct {
	Schedules []HourVote `json:"info" bson:"info"`
}

type HourVote struct {
	Hour  string `json:"hour" bson:"hour"`
	Votes int    `json:"votes" bson:"votes"`
}
