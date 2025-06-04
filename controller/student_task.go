package controller

import (
	"jadwalin/services"
)

type StudentTaskController struct {
	service services.StudentTaskService
}

func NewStudentTaskController(service services.StudentTaskService) *StudentTaskController{
	return &StudentTaskController{
		service: service,
	}
}