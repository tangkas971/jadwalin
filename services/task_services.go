package services

import (
	"fmt"
	"jadwalin/dto"
	"jadwalin/model"
	"jadwalin/repository"
	"jadwalin/utils"
)

type TaskService interface {
	Create(userRole string, input dto.TaskRequestDTO) error
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

func (s *taskService) Create(userRole string, input dto.TaskRequestDTO)error{
	// cek role user
	err := utils.RoleCheck(userRole, "dosen")
	if err != nil {
		return fmt.Errorf("anda adalah seorang %s. hanya dosen yang dapat membuat tugas", userRole)
	}

	task := model.Task{
		Title: input.Title,
		Description: input.Description,
		Deadline: input.Deadline,
		SubjectId: input.SubjectId,
		LecturerId: input.LecturerId,
	}

	err = s.taskRepo.Create(&task)
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

