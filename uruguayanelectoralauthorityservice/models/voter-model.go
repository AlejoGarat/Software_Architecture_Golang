package models

type IdDocumentVoter string
type CredentialVoter = string
type FirstNameVoter = string
type SecondNameVoter = string
type FirstSurnameVoter = string
type SecondSurnameVoter = string
type SexVoter = string
type DateOfBirthVoter = string
type DepartmentResidency = string

type Voter struct {
	IdDocumentVoter     string              `json:"id_document"`
	CredentialVoter     CredentialVoter     `json:"credential"`
	FirstNameVoter      FirstNameVoter      `json:"id_first_name"`
	SecondNameVoter     SecondNameVoter     `json:"id_second_name"`
	FirstSurnameVoter   FirstSurnameVoter   `json:"first_surname"`
	SecondSurnameVoter  SecondSurnameVoter  `json:"second_surname"`
	SexVoter            SexVoter            `json:"id_sex"`
	DateOfBirthVoter    DateOfBirthVoter    `json:"id_date_of_birth"`
	DepartmentResidency DepartmentResidency `json:"department_residency"`
}
