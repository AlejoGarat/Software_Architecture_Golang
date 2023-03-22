package read

type VotersPerAge struct {
	ElectionId string      `json:"election_id" bson:"election_id"`
	CircuitId  string      `json:"circuit_id" bson:"circuit_id"`
	Voters     []AgeAmount `json:"voters_age" bson:"voters_age"`
}

type AgeAmount struct {
	Age          int `json:"age" bson:"age"`
	VotersAmount int `json:"voters" bson:"voters"`
}
