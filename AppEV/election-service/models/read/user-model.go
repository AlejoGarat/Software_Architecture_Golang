package read

type User struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Role     string `json:"role" bson:"role"`
}
