package services

import (
	"jadwalin/dto"
	"jadwalin/model"
	"jadwalin/repository"
)

type TaskService interface {
	Create(input dto.TaskRequestDTO) error
}

type taskService struct {
	taskRepo repository.TaskRepository
	userService UserService
	studentTaskService StudentTaskService
}

func NewTaskService(taskRepo repository.TaskRepository, userService UserService, studentTaskService StudentTaskService) TaskService {
	return &taskService{
		taskRepo: taskRepo,
		userService: userService,
		studentTaskService: studentTaskService,
	}
}


func (s *taskService) Create(input dto.TaskRequestDTO)error{
	task := model.Task{
		Title: input.Title,
		Description: input.Description,
		Deadline: input.Deadline,
		SubjectId: input.SubjectId,
		LecturerId: input.LecturerId,
	}

	err := s.taskRepo.Create(&task)
	if err != nil{
		return err
	}

	// Ambil semua data mahasiswa
	students, err := s.userService.FindByRole("mahasiswa")
	if err != nil {
		return err
	}
	
	// mengambil id setiap students
	var studentIDs []int
	for _, student := range students{
		studentIDs = append(studentIDs, student.Id)
	}

	// Menyebarkan ke semua mahasiswa
	err = s.studentTaskService.AssignTaskToStudents(task.Id, studentIDs)
	if err != nil {
		return err
	}
	
	return nil 
}

