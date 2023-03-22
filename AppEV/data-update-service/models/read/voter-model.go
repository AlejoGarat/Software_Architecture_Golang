package read

type Voter struct {
	Id          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Sex         string `json:"gender" bson:"gender"`
	Surname     string `json:"surname" bson:"surname"`
	DateOfBirth string `json:"birth_date" bson:"birth_date"`
	Circuit     string `json:"circuit_id" bson:"circuit_id"`
	Department  string `json:"voting_department" bson:"voting_department"`
	Cellphone   string `json:"cellphone" bson:"cellphone"`
	Mail        string `json:"mail" bson:"mail"`
	Age         int    `json:"age" bson:"age"`
}
