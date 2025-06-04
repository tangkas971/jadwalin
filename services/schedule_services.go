package services

import (
	"fmt"
	"jadwalin/dto"
	"jadwalin/model"
	"jadwalin/repository"
	"jadwalin/utils"

	"gorm.io/gorm"
)

type ScheduleService interface {
	Create(input dto.ScheduleRequestDTO) error
	GetAll()([]dto.ScheduleResponseDTO, error)
	Delete(id int) error
	Update(id int, input dto.ScheduleRequestDTO) error
}

type scheduleService struct {
	repo repository.ScheduleRepository
	userService UserService
	studentScheduleService StudentSchedulesService
}

func NewScheduleService(repo repository.ScheduleRepository, userService UserService, studentScheduleService StudentSchedulesService) ScheduleService {
	return &scheduleService{
		repo: repo,
		userService: userService,
		studentScheduleService: studentScheduleService,
	}
}

func (s *scheduleService) Create(input dto.ScheduleRequestDTO) error {
	err := utils.RoleCheck(input.UserRole, "dosen")
	if err != nil {
		return fmt.Errorf("anda adalah seorang %s, hanya dosen yang dapat membuat schedules", input.UserRole)
	}

	schedule := model.Schedule{
		Day: input.Day,
		StartTime: input.StartTime,
		EndTime: input.EndTime,
		SubjectId: input.SubjectId,
		LecturerId: input.UserId,
		GradeId: input.GradeId,
		ProdiId: input.ProdiId,
		
	}

	err = s.repo.Create(&schedule)
	if err != nil {
		return  err
	}

	// ambil data student dengan prodi_id dan grade_id tertentu
	students, err := s.userService.FindStudentByGradeAndProdi(input.GradeId, input.ProdiId)
	if err != nil {
		return err
	}

	// mengambil id dari setiap student
	var studentIDs []int
	for _, student := range students{
		studentIDs = append(studentIDs, student.UserId)
	}
	
	// menyebarkan jadwal ke setiap student
	err = s.studentScheduleService.AssignScheduleToStudents(schedule.Id, studentIDs)
	if err != nil {
		return err
	}

	return nil 
	
}

func (s *scheduleService) GetAll()([]dto.ScheduleResponseDTO, error){
	schedules, err := s.repo.GetAll()
	if err != nil {
		return []dto.ScheduleResponseDTO{}, err
	}

	var scheduleDTOs []dto.ScheduleResponseDTO

	for _, schedule := range schedules{
		scheduleDTO := dto.ScheduleResponseDTO{
			Id: schedule.Id,
			Day: schedule.Day,
			StartTime: schedule.StartTime,
			EndTime: schedule.EndTime,
			Subject: dto.SubjectResponseDTO{
				Id: schedule.Subject.Id,
				Name: schedule.Subject.Name,
			},
			Lecturer: dto.LecturerResponseDTO{
				Id: schedule.Lecturer.Id,
				Name: schedule.Lecturer.Name,
			},
			Grade: dto.GradeResponseDTO{
				Id: schedule.Grade.Id,
				Name: schedule.Grade.Name,
			},
			Prodi: dto.ProdiResponseDTO{
				Id: schedule.Prodi.Id,
				Name: schedule.Prodi.Name,
			},
			CreatedAt: schedule.CreatedAt,
			UpdateAt: schedule.UpdatedAt,
		}
		scheduleDTOs = append(scheduleDTOs, scheduleDTO)
	}

	return scheduleDTOs, nil 
}

func (s *scheduleService) Delete(id int) error{
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil 
}

func (s *scheduleService) Update(id int, input dto.ScheduleRequestDTO) error{
	existingSchedule, err := s.repo.FindById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil
		}
		return err
	}

	existingSchedule.Day = input.Day
	existingSchedule.StartTime = input.StartTime
	existingSchedule.EndTime = input.EndTime
	existingSchedule.SubjectId = input.SubjectId
	existingSchedule.GradeId = input.GradeId

	err = s.repo.Update(existingSchedule)
	if err != nil {
		return err
	}

	return nil 
}