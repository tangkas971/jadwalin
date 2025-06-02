package services

import (
	"jadwalin/dto"
	"jadwalin/repository"
)

type TaskService interface {
	FindAll()([]dto.TaskResponseDTO, error)
	// FindById(id int)(dto.TaskResponseDTO, error)
	// Create(input dto.CreateTaskRequest)(dto.TaskResponseDTO, error)
	// Update(id int, input dto.CreateTaskRequest)(dto.TaskResponseDTO, error)
	// Delete(id int) error
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

func (s *taskService) FindAll()([]dto.TaskResponseDTO, error) {
	_, err := s.taskRepo.FindAll()
	if err != nil {
		return []dto.TaskResponseDTO{}, err
	}

	// var taskDTOs []dto.TaskResponseDTO
	// for _, task := range tasks{
	// 	taskDTO := dto.TaskResponseDTO{
	// 		Title: task.Title,
	// 		Description: task.Description,
	// 		// DueDate: task.DueDate,			
	// 		LecturerId: task.LecturerId,
	// 		Lecturer: dto.UserResponseDTO{
	// 			Name: task.Lecturer.Name,
	// 			NIM: task.Lecturer.NIM,
	// 		},
	// 		ScheduleId: task.ScheduleId,
	// 		Schedule: dto.ScheduleResponseDTO{
	// 			Day: task.Schedule.Day,
	// 			SubjectId: task.Schedule.SubjectId,
	// 			Subject: dto.SubjectResponseDTO{
	// 				Code: task.Schedule.Subject.Code,
	// 				Name: task.Schedule.Subject.Name,
	// 			},
	// 		},
	// 	}
	// 	taskDTOs = append(taskDTOs, taskDTO)
	// }

	return []dto.TaskResponseDTO{}, nil 
}

// func (s *taskService) FindById(id int)(dto.TaskResponseDTO, error){
// 	task, err := s.taskRepo.FindById(id)
// 	if err != nil {
// 		return dto.TaskResponseDTO{}, err
// 	}

// 	return dto.TaskResponseDTO{
// 		Title: task.Title,
// 		Description: task.Description,
// 		// DueDate: task.DueDate,
// 		LecturerId: task.LecturerId,
// 		Lecturer: dto.UserResponseDTO{
// 			Name: task.Lecturer.Name,
// 			NIM: task.Lecturer.NIM,
// 		},
// 		ScheduleId: task.ScheduleId,
// 		Schedule: dto.ScheduleResponseDTO{
// 			Day: task.Schedule.Day,
// 			SubjectId: task.Schedule.SubjectId,
// 			Subject: dto.SubjectResponseDTO{
// 				Code: task.Schedule.Subject.Code,
// 				Name: task.Schedule.Subject.Name,
// 			},
// 		},
// 	}, nil 
// }

// func (s *taskService) Create(input dto.CreateTaskRequest)(dto.TaskResponseDTO, error){
// 	task := model.Task{
// 		Title: input.Title,
// 		Description: input.Description,
// 		Deadline: input.Deadline,
// 		LecturerId: input.LecturerId,
// 		ScheduleId: input.ScheduleId,
// 	}

// 	err := s.taskRepo.Create(&task)
// 	if err != nil{
// 		return dto.TaskResponseDTO{}, err
// 	}

// 	// Ambil semua data mahasiswa
// 	students, err := s.userService.FindByRole("mahasiswa")
// 	if err != nil {
// 		return dto.TaskResponseDTO{}, err
// 	}
	
// 	// mengambil id setiap students
// 	var studentIDs []int
// 	for _, student := range students{
// 		studentIDs = append(studentIDs, student.Id)
// 	}

// 	// Menyebarkan ke semua mahasiswa
// 	err = s.studentTaskService.AssignTaskToStudents(task.Id, studentIDs)
// 	if err != nil {
// 		return dto.TaskResponseDTO{}, err
// 	}

// 	return dto.TaskResponseDTO{
// 		Title: task.Title,
// 		Description: task.Description,
// 		LecturerId: task.LecturerId,
// 		Lecturer: dto.UserResponseDTO{
// 			Name: task.Lecturer.Name,
// 		},
// 		ScheduleId: task.ScheduleId,
// 		Schedule: dto.ScheduleResponseDTO{
// 			Day: task.Schedule.Day,
// 			SubjectId: task.Schedule.SubjectId,
// 			Subject: dto.SubjectResponseDTO{
// 				Code: task.Schedule.Subject.Code,
// 				Name: task.Schedule.Subject.Name,
// 			},
// 		},
// 	}, nil  
// }
