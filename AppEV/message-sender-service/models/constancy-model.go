package models

type Info struct {
	Date string `bson:"date"`
	Hour string `bson:"hour"`
}

type Constancy struct {
	Election             Election `json:"election_data"`
	VoteEmissionInfoDate Info     `json:"vote_emission_date"`
	VoterDocument        string   `json:"voter_document"`
	Name                 string   `json:"voter_name"`
	Surname              string   `json:"voter_surname"`
	VoteId               string   `json:"vote_id"`
}

type ConstancyDBData struct {
	VoterDocument string `bson:"voter_document"`
	Timestamp     int    `bson:"timestamp"`
}
