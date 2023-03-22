package write

import (
	"time"
)

type Election struct {
	Id              string    `json:"id" bson:"id"`
	Description     string    `json:"description" bson:"description"`
	Url             string    `json:"url" bson:"url"`
	StartDate       time.Time `json:"start_date" bson:"start_date"`
	EndDate         time.Time `json:"end_date" bson:"end_date"`
	VotationMode    string    `json:"votation_mode" bson:"votation_mode"`
	EligibleVoters  int       `json:"eligible_voters" bson:"eligible_voters"`
	TotalVoteAmount int       `json:"total_votes" bson:"total_votes"`
}
