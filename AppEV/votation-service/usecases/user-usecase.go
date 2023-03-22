package usecases

import (
	"errors"
	idataaccess "votation-service/dataaccess/interfaces"
	"votation-service/helpers"
	"votation-service/models/write"
)

type UserUseCase struct {
	userRepository idataaccess.UserRepository
	helpers        helpers.Helpers
}

func NewUserUseCase(userRepository idataaccess.UserRepository, helpers helpers.Helpers) *UserUseCase {
	return &UserUseCase{userRepository: userRepository, helpers: helpers}
}

func (userUseCase *UserUseCase) Login(username string, password string) error {
	var log write.LoggingModel

	if username != "john" || password != "doe" {
		log = write.LoggingModel{Type: "Error", Operation: "Login", Actor: "Voter", Description: "Error in parameters"}
		userUseCase.helpers.LogHelper.SendLog(log)

		return errors.New(" Usted no posee autorización")
	}

	return nil

	/*user, err := userUseCase.userRepository.FindByUsername(username)

	// hasheamos password y comparamos con lo guardado en la db

	if user.Password == password {
		return errors.New(" Credenciales incorrectas")
	}

	if err == nil {
		return err
	}

	return err*/
}

func (userUseCase UserUseCase) Logout() error {
	return nil
}
