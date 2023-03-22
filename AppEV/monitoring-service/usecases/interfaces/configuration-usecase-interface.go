package interfaces

import (
	"monitoring-service/rabbit/workers"
)

type ConfigurationUseCase interface {
	AnalyzeValues(workers.Worker) ([]string, error)
}
