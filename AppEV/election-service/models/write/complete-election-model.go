package write

import (
	"election-service/models/read"
	"time"
)

type CompleteElection struct {
	Id               string           `json:"id" bson:"id"`
	Description      string           `json:"description" bson:"description"`
	Url              string           `json:"url" bson:"url"`
	Circuits         []Circuit        `json:"circuits" bson:"circuits"`
	Candidates       []Candidate      `json:"candidates" bson:"candidates"`
	Voters           []read.Voter     `json:"voters" bson:"voters"`
	PoliticalParties []PoliticalParty `json:"political_parties" bson:"political_parties"`
	StartDate        time.Time        `json:"start_date" bson:"start_date"`
	EndDate          time.Time        `json:"end_date" bson:"end_date"`
	VotationMode     string           `json:"votation_mode" bson:"votation_mode"`
}
