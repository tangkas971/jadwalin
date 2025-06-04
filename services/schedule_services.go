package services

import (
	"fmt"
	"jadwalin/dto"
	"jadwalin/model"
	"jadwalin/repository"
	"jadwalin/utils"
)

type ScheduleService interface {
	Create(input dto.ScheduleRequestDTO) error
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
