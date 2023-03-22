package datasources

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func NewRedisDataSource(connectionURI string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{Addr: connectionURI})

	err := client.Ping(context.Background()).Err()

	if err != nil {
		return client, err
	}

	return client, nil
}
