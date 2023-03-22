package write

type Candidate struct {
	ElectionId       string `json:"election_id" bson:"election_id"`
	CandidateId      string `json:"id" bson:"id"`
	PoliticalPartyId string `json:"political_party_id" bson:"political_party_id"`
	Name             string `json:"name" bson:"name"`
	Surname          string `json:"surname" bson:"surname"`
	Gender           string `json:"gender" bson:"gender"`
	BirthDate        string `json:"birth_date" bson:"birth_date"`
}
