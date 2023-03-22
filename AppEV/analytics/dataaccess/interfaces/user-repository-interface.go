package interfaces

type UserRepository interface {
	GetUserRole(string) (string, error)
}
