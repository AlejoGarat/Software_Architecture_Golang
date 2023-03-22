package common

import (
	"analytics/api/controllers"
)

type Controllers struct {
	CircuitController    controllers.CircuitController
	DepartmentController controllers.DepartmentController
	ScheduleController   controllers.ScheduleController
}
