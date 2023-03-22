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
	Id          Id          `json:"id"`
	Name        Name        `json:"name"`
	Surname     Surname     `json:"surname" `
	Sex         Sex         `json:"sex" `
	DateOfBirth DateOfBirth `json:"birth_date" `
	Circuit     Circuit     `json:"circuit_id" `
	Department  Department  `json:"voting_department" `
	Cellphone   Cellphone   `json:"cellphone" `
	Mail        Mail        `json:"mail" `
}
