package interfaces

type UserUseCase interface {
	Login(string, string) (string, error)
	HashPassword(password string) string
}
