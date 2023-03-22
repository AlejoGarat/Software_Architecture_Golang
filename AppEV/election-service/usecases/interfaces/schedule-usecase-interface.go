package interfaces

import "election-service/models/read"

type ScheduleUseCase interface {
	GetFrequentSchedules(electionId string) (read.FrequentVotationSchedules, error)
}
