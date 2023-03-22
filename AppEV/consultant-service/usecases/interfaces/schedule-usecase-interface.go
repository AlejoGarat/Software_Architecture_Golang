package interfaces

import "consultant-service/models/read"

type ScheduleUseCase interface {
	GetFrequentSchedules(electionId string) (read.FrequentVotationSchedules, error)
}
