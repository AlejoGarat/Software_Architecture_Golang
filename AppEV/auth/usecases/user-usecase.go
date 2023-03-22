package usecases

import (
	idataaccess "auth/dataaccess/interfaces"
	"auth/helpers"
	"auth/models/read"
	"auth/models/write"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserUseCase struct {
	userRepository idataaccess.UserRepository
	helpers        helpers.Helpers
}

func NewUserUseCase(userRepository idataaccess.UserRepository, helpers helpers.Helpers) *UserUseCase {
	return &UserUseCase{userRepository: userRepository, helpers: helpers}
}

func (userUseCase *UserUseCase) Login(id string, password string) (string, error) {
	var log write.LoggingModel

	var token string

	_, err := userUseCase.userRepository.FindUser(id, password)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Login", Actor: "User", Description: "Error in parameters"}
		userUseCase.helpers.LogHelper.SendLog(log)

		return token, errors.New("incorrect credentials")
	}

	var user read.User

	user.Id = id
	user.Password = password
	token, err = userUseCase.GetToken(user)

	return token, err
}

func (userUseCase *UserUseCase) HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (userUseCase *UserUseCase) GetToken(user read.User) (string, error) {
	claims := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	strToken, err := token.SignedString([]byte("secret"))

	if err != nil {
		return strToken, err
	}

	err = userUseCase.userRepository.AddTokenToUser(user.Id, strToken)

	if err != nil {
		return strToken, err
	}

	return strToken, nil
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjE5NTcxMzZ9.RB3arc4-OyzASAaUhC2W3ReWaXAt_z2Fd3BN4aWTgEY
