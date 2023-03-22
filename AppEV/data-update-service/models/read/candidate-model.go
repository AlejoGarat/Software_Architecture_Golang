package read

type Candidate struct {
	CandidateId      string `json:"id" bson:"id"`
	PoliticalPartyId string `json:"political_party_id" bson:"political_party_id"`
}
