package interfaces

import "analytics/models/read"

type ScheduleUseCase interface {
	GetFrequentSchedules(electionId string) (read.FrequentVotationSchedules, error)
}
