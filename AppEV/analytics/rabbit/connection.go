package rabbit

import (
	workers "analytics/rabbit/workers"
)

type RabbitConnection struct{}

func NewRabbitConnection() (workers.Worker, error) {
	aworker, err := workers.BuildRabbitWorker("amqp://guest:guest@localhost:5672/")

	if err != nil {
		return nil, err
	}

	return aworker, nil
}
