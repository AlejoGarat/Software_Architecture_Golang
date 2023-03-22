package common

import (
	"election-service/api/controllers"
)

type Controllers struct {
	ElectionController   controllers.ElectionController
	CircuitController    controllers.CircuitController
	DepartmentController controllers.DepartmentController
	ScheduleController   controllers.ScheduleController
	AlertController      controllers.AlertController
	VoteController       controllers.VoteController
}
