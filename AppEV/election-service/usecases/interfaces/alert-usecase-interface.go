package interfaces

import "election-service/models/read"

type AlertUseCase interface {
	GetAlertConfiguration(electionId string) (read.AlertConfiguration, error)
}
