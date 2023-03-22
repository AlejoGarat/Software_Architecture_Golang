package models

type ElectionId = string

type Election struct {
	ElectionId  ElectionId `json:"id" bson:"id"`
	Description string     `json:"description" bson:"id"`
	Start       string     `json:"start_date" bson:"start_date"`
	End         string     `json:"end_date" bson:"end_date"`
	Votes       int        `json:"votes" bson:"votes"`
}
