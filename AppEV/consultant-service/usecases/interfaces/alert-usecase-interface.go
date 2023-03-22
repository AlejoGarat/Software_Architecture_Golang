package interfaces

import (
	"consultant-service/models/write"
)

type AlertUseCase interface {
	ModifyAlertConfiguration(write.AlertConfiguration) error
	GetAlertConfiguration(electionId string) (write.AlertConfiguration, error)
}
