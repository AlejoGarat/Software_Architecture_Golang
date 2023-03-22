package interfaces

type UserUseCase interface {
	Login(email string, password string) error
	Logout() error
}
