package interfaces

import (
	"auth/models/read"
)

type UserRepository interface {
	FindUser(string, string) (read.User, error)
	AddTokenToUser(id string, token string) error
}
