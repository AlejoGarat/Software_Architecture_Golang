package interfaces

import (
	"election-service/models/read"
)

type ScheduleRepository interface {
	GetFrequentSchedules(electionId string) (read.FrequentVotationSchedules, error)
}
