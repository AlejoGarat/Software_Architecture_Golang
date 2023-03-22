package factory

import (
	"consultant-service/api/controllers"
	"consultant-service/dataaccess"
	idataaccess "consultant-service/dataaccess/interfaces"
	"consultant-service/helpers"
	"consultant-service/usecases"
	iusecases "consultant-service/usecases/interfaces"

	"github.com/go-redis/redis/v8"
)

type AlertFactory struct {
	redisCli *redis.Client
	helpers  helpers.Helpers
}

func NewAlertFactory(redisCli *redis.Client, helpers helpers.Helpers) *AlertFactory {
	return &AlertFactory{redisCli: redisCli, helpers: helpers}
}

func (AlertFactory *AlertFactory) GetAlertController() *controllers.AlertController {
	AlertRepository := InitializeAlertRepository(AlertFactory.redisCli)
	AlertUseCase := InitializeAlertUseCase(AlertRepository, AlertFactory.helpers)
	AlertController := InitializeAlertController(AlertUseCase)
	return AlertController
}

func InitializeAlertRepository(redisCli *redis.Client) idataaccess.AlertRepository {
	AlertRepo := dataaccess.NewAlertRepo(redisCli)
	var AlertRepository idataaccess.AlertRepository = AlertRepo
	return AlertRepository
}

func InitializeAlertUseCase(AlertRepository idataaccess.AlertRepository, helpers helpers.Helpers) *usecases.AlertUseCase {
	return usecases.NewAlertUseCase(AlertRepository, helpers)
}

func InitializeAlertController(AlertUseCase iusecases.AlertUseCase) *controllers.AlertController {
	return controllers.NewAlertController(AlertUseCase)
}
