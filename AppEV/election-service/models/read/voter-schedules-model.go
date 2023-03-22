package read

type VoterSchedules struct {
	Info []Info `json:"info"`
}

type VoterSchedulesRequest struct {
	Election string `json:"election_id" bson:"election_id"`
	VoterId  string `json:"voter_id" bson:"voter_id"`
}
