package read

type VotersPerDepartment struct {
	ElectionId   string `json:"election_id" bson:"election_id"`
	Department   string `json:"department" bson:"department"`
	VotersAmount int    `json:"voters" bson:"voters"`
}

type VotersPerDepartmentByGenderAndAge struct {
	ElectionId      string         `json:"election_id" bson:"election_id"`
	Department      string         `json:"department" bson:"department"`
	VotersPerAge    []AgeAmount    `json:"voters_age" bson:"voters_age"`
	VotersPerGender []GenderAmount `json:"voters_gender" bson:"voters_gender"`
}
