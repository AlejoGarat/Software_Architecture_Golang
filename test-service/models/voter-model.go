package models

type Voter struct {
	Id                  string `json:"id" bson:"id"`
	Credential          string `json:"credential" bson:"credential"`
	Name                string `json:"name" bson:"name"`
	Surname             string `json:"surname" bson:"surname"`
	Gender              string `json:"gender" bson:"gender"`
	BirthDate           string `json:"birth_date" bson:"birth_date"`
	CircuitId           string `json:"circuit_id" bson:"circuit_id"`
	ResidenceDepartment string `json:"residence_department" bson:"residence_department"`
	Celphone            string `json:"cellphone" bson:"cellphone"`
	MailAddress         string `json:"mail" bson:"mail"`
	VotingDepartment    string `json:"voting_department" bson:"voting_department"`
}
