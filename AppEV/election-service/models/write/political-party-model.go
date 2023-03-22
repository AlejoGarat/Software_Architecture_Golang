package write

type PoliticalParty struct {
	Id         string `json:"id" bson:"id"`
	Name       string `json:"name" bson:"name"`
	ElectionId string `json:"election_id" bson:"election_id"`
}
