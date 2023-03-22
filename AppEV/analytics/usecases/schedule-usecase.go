package usecases

import (
	idataaccess "analytics/dataaccess/interfaces"
	"analytics/helpers"
	"analytics/models/read"
	"analytics/models/write"
)

type ScheduleUseCase struct {
	scheduleRepository idataaccess.ScheduleRepository
	helpers            helpers.Helpers
}

func NewScheduleUseCase(scheduleRepository idataaccess.ScheduleRepository, helpers helpers.Helpers) *ScheduleUseCase {
	return &ScheduleUseCase{scheduleRepository: scheduleRepository, helpers: helpers}
}

func (scheduleUseCase ScheduleUseCase) GetFrequentSchedules(electionId string) (read.FrequentVotationSchedules, error) {
	var log write.LoggingModel
	frequentSchedules, err := scheduleUseCase.scheduleRepository.GetFrequentSchedules(electionId)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Get Votation Frequent Schedules", Actor: "Consultant", Description: err.Error()}
		scheduleUseCase.helpers.LogHelper.SendLog(log)
	}

	return frequentSchedules, err
}
