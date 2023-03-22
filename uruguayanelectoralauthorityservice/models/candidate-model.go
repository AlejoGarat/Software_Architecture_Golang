package models

type IdDocument = string
type Credential = string
type FirstName = string
type SecondName = string
type FirstSurname = string
type SecondSurname = string
type Sex = string
type DateOfBirth = string
type IdParty = string

type Candidate struct {
	IdDocument    IdDocument    `json:"id_document"`
	FirstName     FirstName     `json:"id_first_name"`
	SecondName    SecondName    `json:"id_second_name"`
	FirstSurname  FirstSurname  `json:"first_surname"`
	SecondSurname SecondSurname `json:"second_surname"`
	Sex           Sex           `json:"id_sex"`
	DateOfBirth   DateOfBirth   `json:"id_date_of_birth"`
	IdParty       IdParty       `json:"id_party"`
}
