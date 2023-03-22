package factory

import (
	"votation-service/dataaccess"
	idataaccess "votation-service/dataaccess/interfaces"

	"github.com/go-redis/redis/v8"
)

type FilterFactory struct {
	redisCli *redis.Client
}

func NewFilterFactory(redisCli *redis.Client) *FilterFactory {
	return &FilterFactory{redisCli: redisCli}
}

func (FilterFactory *FilterFactory) InitializeFilterRepository(redisCli *redis.Client) idataaccess.FiltersRepository {
	FilterRepo := dataaccess.NewFilterRepository(redisCli)
	var FilterRepository idataaccess.FiltersRepository = FilterRepo
	return FilterRepository
}
