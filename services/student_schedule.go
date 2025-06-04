package services

import (
	"jadwalin/model"
	"jadwalin/repository"
)

type StudentSchedulesService interface {
	AssignScheduleToStudents(scheduleId int, studentIds []int) error
}

type studentScheduleService struct {
	repo repository.StudentScheduleRepository
}

func NewStudentSchedulesService(repo repository.StudentScheduleRepository) StudentSchedulesService{
	return &studentScheduleService{
		repo: repo,
	}
}

func (r *studentScheduleService) AssignScheduleToStudents(scheduleId int, studentIds []int) error{
	for _, studentId := range studentIds{
		StudentSchedule := model.StudentSchedules{
			StudentId: studentId,
			ScheduleId: scheduleId,
		}
		err := r.repo.Create(&StudentSchedule)
		if err != nil {
			return err
		}
	}
	return nil 
}