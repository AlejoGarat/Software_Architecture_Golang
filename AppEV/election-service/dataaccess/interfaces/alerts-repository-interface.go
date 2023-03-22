package interfaces

import (
	"election-service/models/read"
)

type AlertRepository interface {
	GetAlertConfiguration(electionId string) (read.AlertConfiguration, error)
}
