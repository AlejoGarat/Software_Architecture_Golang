package read

type Username = string
type Password = string

type User struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}
