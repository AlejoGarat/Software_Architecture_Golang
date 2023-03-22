package usecases

import (
	idataaccess "consultant-service/dataaccess/interfaces"
	"consultant-service/helpers"
	"consultant-service/models/read"
	"consultant-service/models/write"
)

type FilterUseCase struct {
	filterRepository idataaccess.FilterRepository
	helpers          helpers.Helpers
}

func NewFilterUseCase(filterRepository idataaccess.FilterRepository, helpers helpers.Helpers) *FilterUseCase {
	return &FilterUseCase{filterRepository: filterRepository, helpers: helpers}
}

func (filterUseCase *FilterUseCase) ModifyElectionBeginningFilters(filters read.ElectionBeginningFilters) error {
	var log write.LoggingModel
	err := filterUseCase.filterRepository.ModifyElectionBeginningFilters(filters)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Modify Election Beginning Filters", Actor: "Consultant", Description: err.Error()}
		filterUseCase.helpers.LogHelper.SendLog(log)
	}

	return err
}

func (filterUseCase *FilterUseCase) ModifyElectionEndFilters(filters read.ElectionEndFilters) error {
	var log write.LoggingModel
	err := filterUseCase.filterRepository.ModifyElectionEndFilters(filters)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Modify Election Beginning Filters", Actor: "Consultant", Description: err.Error()}
		filterUseCase.helpers.LogHelper.SendLog(log)
	}

	return err
}

func (filterUseCase *FilterUseCase) ModifyVoteIssuanceFilters(filters read.VoteIssuanceFilters) error {
	var log write.LoggingModel
	err := filterUseCase.filterRepository.ModifyVoteIssuanceFilters(filters)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Modify Election Beginning Filters", Actor: "Consultant", Description: err.Error()}
		filterUseCase.helpers.LogHelper.SendLog(log)
	}

	return err
}
