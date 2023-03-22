package read

type ElectionData struct {
	Id          string `json:"id" bson:"id"`
	Description string `json:"description" bson:"description"`
	Url         string `json:"url" bson:"url"`
	StartDate   string `json:"start_date" bson:"start_date"`
	StartTime   string `json:"start_time" bson:"start_time"`
}
