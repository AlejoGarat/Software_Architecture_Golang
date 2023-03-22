package interfaces

import (
	"votation-service/models/write"
)

type Messenger interface {
	SendMessage(constancy write.Constancy) error
}
