package read

import "time"

type Election struct {
	ElectionId string    `json:"id"`
	EndDate    time.Time `json:"end_date"`
}

type CompleteElection struct {
	Id           string    `json:"id"`
	Descritption string    `json:"description"`
	Url          string    `json:"url"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	VotationMode string    `json:"votation_mode"`
	Voters       int       `json:"eligible_voters"`
	Votes        int       `json:"total_votes"`
}
