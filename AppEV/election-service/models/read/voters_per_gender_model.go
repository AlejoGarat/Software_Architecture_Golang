package read

type VotersPerGender struct {
	ElectionId string         `json:"election_id" bson:"election_id"`
	CircuitId  string         `json:"circuit_id" bson:"circuit_id"`
	Voters     []GenderAmount `json:"voters_gender" bson:"voters_gender"`
}

type GenderAmount struct {
	Gender       string `json:"gender" bson:"gender"`
	VotersAmount int    `json:"voters" bson:"voters"`
}
