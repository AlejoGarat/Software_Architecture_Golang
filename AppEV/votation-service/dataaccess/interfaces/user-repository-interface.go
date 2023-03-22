package interfaces

type UserRepository interface {
	GetUserRole(string, string) (string, error)
}
