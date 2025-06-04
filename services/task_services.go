package services

import (
	"jadwalin/dto"
	"jadwalin/model"
	"jadwalin/repository"
	"jadwalin/utils"
)

type TaskService interface {
	Create(userRole string, input dto.TaskRequestDTO) error
	GetAll()([]dto.TaskResponseDTO, error)
	Delete(id int) error
	Update(id int, input dto.TaskRequestDTO) error
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
		return err
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

func (s *taskService) GetAll()([]dto.TaskResponseDTO, error){
	tasks, err := s.taskRepo.GetAll()
	if err != nil {
		return []dto.TaskResponseDTO{}, err
	}

	var taskDTOs []dto.TaskResponseDTO

	for _, task := range tasks{
		taskDTO := dto.TaskResponseDTO{
			Id: task.Id,
			Title: task.Title,
			Description: task.Description,
			Deadline: task.Deadline,
			Subject: dto.SubjectResponseDTO{
				Id: task.Subject.Id,
				Name: task.Subject.Name,
			},
			Lecturer: dto.LecturerResponseDTO{
				Id: task.Lecturer.Id,
				Name: task.Lecturer.Name,
			},
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		}
		taskDTOs = append(taskDTOs, taskDTO)
	}

	return taskDTOs, nil 
}

func (s *taskService) Delete(id int) error{
	err := s.taskRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil 
}

func (s *taskService) Update(id int, input dto.TaskRequestDTO) error{
	existingTask, err := s.taskRepo.FindById(id)
	if err != nil {
		return err
	}

	existingTask.Title = input.Title
	existingTask.Description = input.Description
	existingTask.Deadline = input.Deadline
	existingTask.SubjectId = input.SubjectId

	err = s.taskRepo.Update(existingTask)
	if err != nil {
		return err
	}

	return nil
}