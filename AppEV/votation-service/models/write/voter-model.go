package write

type Id = string
type Name = string
type Surname = string
type Sex = string
type DateOfBirth = string
type Circuit = string
type Department = string
type Cellphone = string
type Mail = string

type Voter struct {
	Id          Id          `json:"id" bson:"id"`
	Name        Name        `json:"name" bson:"name"`
	Sex         Sex         `json:"gender" bson:"gender"`
	Surname     Surname     `json:"surname" bson:"surname"`
	DateOfBirth DateOfBirth `json:"birth_date" bson:"birth_date"`
	Circuit     Circuit     `json:"circuit_id" bson:"circuit_id"`
	Department  Department  `json:"voting_department" bson:"voting_department"`
	Cellphone   Cellphone   `json:"cellphone" bson:"cellphone"`
	Mail        Mail        `json:"mail" bson:"mail"`
}
