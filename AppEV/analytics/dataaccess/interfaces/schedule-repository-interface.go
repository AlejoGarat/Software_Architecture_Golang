package interfaces

import (
	"analytics/models/read"
)

type ScheduleRepository interface {
	GetFrequentSchedules(electionId string) (read.FrequentVotationSchedules, error)
}
