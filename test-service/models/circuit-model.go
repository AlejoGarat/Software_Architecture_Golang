package models

type Circuit struct {
	Id         string `json:"id" bson:"id"`
	ElectionId string `json:"election_id" bson:"election_id"`
	Department string `json:"department" bson:"deparment"`
	Location   string `json:"location" bson:"location"`
}
