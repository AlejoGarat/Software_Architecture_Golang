package interfaces

import (
	"consultant-service/models/read"
)

type ScheduleRepository interface {
	GetFrequentSchedules(electionId string) (read.FrequentVotationSchedules, error)
}
