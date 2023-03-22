package interfaces

import "consultant-service/models/write"

type AlertRepository interface {
	ModifyAlertConfiguration(write.AlertConfiguration) error
	GetAlertConfiguration(electionId string) (write.AlertConfiguration, error)
}
