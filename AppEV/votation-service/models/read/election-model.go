package read

import "time"

type Election struct {
	ElectionId   ElectionId `json:"id" bson:"id"`
	Description  string     `json:"description" bson:"id"`
	Start        time.Time  `json:"start_date" bson:"start_date"`
	End          time.Time  `json:"end_date" bson:"end_date"`
	Votes        int        `json:"votes" bson:"votes"`
	VotationMode string     `json:"votation_mode" bson:"votation_mode"`
}
