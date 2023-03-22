package write

import (
	"time"
	"votation-service/models/read"
)

type Constancy struct {
	Election             read.Election `json:"election_data"`
	VoteEmissionInfoDate Info          `json:"vote_emission_date"`
	VoterDocument        string        `json:"voter_document"`
	Name                 string        `json:"voter_name"`
	Surname              string        `json:"voter_surname"`
	VoteId               string        `json:"vote_id"`
}

type ConstancyDBData struct {
	VoterDocument string    `bson:"voter_document"`
	Timestamp     time.Time `bson:"timestamp"`
}
