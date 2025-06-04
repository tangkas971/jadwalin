package services

import (
	"jadwalin/model"
	"jadwalin/repository"
)

type StudentTaskService interface {
	AssignTaskToStudents(taskId int, studentIDs []int) error
}

type studentTaskService struct {
	repo repository.StudentTaskRepository
}

func NewStudentTaskService(repo repository.StudentTaskRepository) StudentTaskService {
	return &studentTaskService{
		repo: repo,
	}
}

// menyebarkan satu tugas ke banyak mahasiswa
func (s *studentTaskService) AssignTaskToStudents(taskId int, studentIDs []int) error {
	for _, studentID := range studentIDs {
		studentTask := &model.StudentTask{
			StudentId: studentID,
			TaskId: taskId,
			Status: "pending",
		}
		err := s.repo.Create(studentTask)
		if err != nil {
			return err
		}
	}

	return nil 
}
