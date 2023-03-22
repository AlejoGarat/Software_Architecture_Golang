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

type FilterFactory struct {
	redisCli *redis.Client
	helpers  helpers.Helpers
}

func NewFilterFactory(redisCli *redis.Client, helpers helpers.Helpers) *FilterFactory {
	return &FilterFactory{redisCli: redisCli, helpers: helpers}
}

func (FilterFactory *FilterFactory) GetFilterController() *controllers.FilterController {
	FilterRepository := InitializeFilterRepository(FilterFactory.redisCli)
	FilterUseCase := InitializeFilterUseCase(FilterRepository, FilterFactory.helpers)
	FilterController := InitializeFilterController(FilterUseCase)
	return FilterController
}

func InitializeFilterRepository(redisCli *redis.Client) idataaccess.FilterRepository {
	FilterRepo := dataaccess.NewFilterRepo(redisCli)
	var FilterRepository idataaccess.FilterRepository = FilterRepo
	return FilterRepository
}

func InitializeFilterUseCase(FilterRepository idataaccess.FilterRepository, helpers helpers.Helpers) *usecases.FilterUseCase {
	return usecases.NewFilterUseCase(FilterRepository, helpers)
}

func InitializeFilterController(FilterUseCase iusecases.FilterUseCase) *controllers.FilterController {
	return controllers.NewFilterController(FilterUseCase)
}
