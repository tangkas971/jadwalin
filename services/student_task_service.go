package services

import (
	"jadwalin/dto"
	"jadwalin/model"
	"jadwalin/repository"
)

type StudentTaskService interface {
	AssignTaskToStudents(taskId int, studentIDs []int) error
	FindByStudentTaskId(student_id uint)([]dto.StudentTaskResponse, error)
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

func (s *studentTaskService) FindByStudentTaskId(student_id uint)([]dto.StudentTaskResponse, error){
	studentTasks, err := s.repo.FindByStudentId(student_id)
	if err != nil {
		return []dto.StudentTaskResponse{}, err
	}

	var studentTaskDTOs []dto.StudentTaskResponse
	for _,studentTask := range studentTasks {
		studentTaskDTO := dto.StudentTaskResponse{
			Id: studentTask.Id,
			TaskId: studentTask.TaskId,
			Task: dto.TaskResponseDTO{
				Title: studentTask.Task.Title,
				Description: studentTask.Task.Description,
				LecturerId: studentTask.Task.LecturerId,
				Lecturer: dto.UserResponseDTO{
					Id: studentTask.Task.Lecturer.Id,
					Name: studentTask.Task.Lecturer.Name,
				},
				ScheduleId: studentTask.Task.ScheduleId,
				Schedule: dto.ScheduleResponseDTO{
					Id: studentTask.Task.Schedule.Id,
					Day: studentTask.Task.Schedule.Day,
					SubjectId: studentTask.Task.Schedule.SubjectId,
					Subject: dto.SubjectResponseDTO{
						Id: studentTask.Task.Schedule.Subject.Id,
						Code: studentTask.Task.Schedule.Subject.Code,
						Name: studentTask.Task.Schedule.Subject.Name,
					},
				},
			},
			Status: studentTask.Status,
			SubmittedAt: studentTask.SubmittedAt,
		}
		studentTaskDTOs = append(studentTaskDTOs, studentTaskDTO)
	}

	return studentTaskDTOs, nil 
}