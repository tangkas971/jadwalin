package services

import (
	"fmt"
	"jadwalin/dto"
	"jadwalin/model"
	"jadwalin/repository"

	"gorm.io/gorm"
)

type ScheduleService interface {
	Create(idUser uint, roleUser string, input dto.CreateScheduleRequestDTO)(dto.ScheduleResponseDTO, error)
	FindAll()([]dto.ScheduleResponseDTO, error)
	FindById(id int)(dto.ScheduleResponseDTO, error)
	Update(roleUser string, id int, input dto.UpdateScheduleRequestDTO)(dto.ScheduleResponseDTO, error)
	Delete(roleUser string, id int) error
}

type scheduleService struct {
	repo repository.ScheduleRepository
}

func NewScheduleService(repo repository.ScheduleRepository) ScheduleService {
	return &scheduleService{repo: repo}
}

func (s *scheduleService) Create(idUser uint, roleUser string, input dto.CreateScheduleRequestDTO)(dto.ScheduleResponseDTO, error){
	if roleUser != "dosen" {
		return dto.ScheduleResponseDTO{}, fmt.Errorf("hanya dosen yang dapat membuat schedule")
	}

	schedule := model.Schedule{
		Day: input.Day,
		StartTime: input.StartTime,
		EndTime: input.EndTime,
		SubjectId: input.SubjectId,
		LecturerId: int(idUser),
	}

	err := s.repo.Create(&schedule)
	if err != nil {
		return dto.ScheduleResponseDTO{}, err
	}

	return dto.ScheduleResponseDTO{
		Id: schedule.Id,
		Day: schedule.Day,
		StartTime: schedule.StartTime,
		EndTime: schedule.EndTime,
		SubjectId: schedule.SubjectId,
		LecturerId: int(idUser),
	}, nil 
	
}

func (s *scheduleService) Delete(roleUser string, id int) error {
	if roleUser != "dosen" {
		return fmt.Errorf("hanya dosen yang dapat menghapus schedule")
	}
	_, err := s.repo.FindById(id)
	if err != nil {
		return fmt.Errorf("schedule dengan id %d tidak ditemukan: %w", id, err)
	}

	err = s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
	
}

func (s *scheduleService) FindAll()([]dto.ScheduleResponseDTO, error){
	schedules, err := s.repo.FindAll()
	if err != nil {
		return []dto.ScheduleResponseDTO{}, err
	}

	var scheduleDTOs []dto.ScheduleResponseDTO
	for _, schedule := range schedules {
		scheduleDTO := dto.ScheduleResponseDTO{
			Id: schedule.Id,
			Day: schedule.Day,
			StartTime: schedule.StartTime,
			EndTime: schedule.EndTime,
			SubjectId: schedule.SubjectId,
			Subject: dto.SubjectResponseDTO{
				Id: schedule.Subject.Id,
				Code: schedule.Subject.Code,
				Name: schedule.Subject.Name,
			},
			LecturerId: schedule.LecturerId,
			Lecturer: dto.LecturerResponseDTO{
				Id: schedule.Lecturer.Id,
				Name: schedule.Lecturer.Name,
				Email: schedule.Lecturer.Email,
			},
		}
		scheduleDTOs = append(scheduleDTOs, scheduleDTO)
	}

	return scheduleDTOs, nil 
}

func (s *scheduleService) FindById(id int)(dto.ScheduleResponseDTO, error){
	schedule, err := s.repo.FindById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return dto.ScheduleResponseDTO{}, fmt.Errorf("schedule dengan id %d tidak ditemukan", id)
		}
		return dto.ScheduleResponseDTO{}, err
	}

	return dto.ScheduleResponseDTO{
		Id: schedule.Id,
		Day: schedule.Day,
		StartTime: schedule.StartTime,
		EndTime: schedule.EndTime,
		SubjectId: schedule.SubjectId,
		Subject: dto.SubjectResponseDTO{
			Id: schedule.Subject.Id,
			Code: schedule.Subject.Code,
			Name: schedule.Subject.Name,
		},
		LecturerId: schedule.LecturerId,
		Lecturer: dto.LecturerResponseDTO{
			Id: schedule.Lecturer.Id,
			Name: schedule.Lecturer.Name,
			Email: schedule.Lecturer.Email,
		},
	}, nil 
}

func (s *scheduleService) Update(roleUser string, id int, input dto.UpdateScheduleRequestDTO)(dto.ScheduleResponseDTO, error){
	if roleUser != "dosen" {
		return dto.ScheduleResponseDTO{}, fmt.Errorf("hanya dosen yang dapat mengupdate schedule")
	}
	existingSchedule, err := s.repo.FindById(id)
	if err != nil {
		return dto.ScheduleResponseDTO{}, err
	}

	existingSchedule.Day = input.Day
	existingSchedule.StartTime = input.StartTime
	existingSchedule.EndTime = input.EndTime

	err = s.repo.Update(existingSchedule)
	if err != nil {
		return dto.ScheduleResponseDTO{}, err
	}

	return dto.ScheduleResponseDTO{
		Id: existingSchedule.Id,
		Day: existingSchedule.Day,
		StartTime: existingSchedule.StartTime,
		EndTime: existingSchedule.EndTime,
		SubjectId: existingSchedule.SubjectId,
		Subject: dto.SubjectResponseDTO{
			Id: existingSchedule.Subject.Id,
			Code: existingSchedule.Subject.Code,
			Name: existingSchedule.Subject.Name,
		},
		LecturerId: existingSchedule.LecturerId,
		Lecturer: dto.LecturerResponseDTO{
			Id: existingSchedule.Lecturer.Id,
			Name: existingSchedule.Lecturer.Name,
			Email: existingSchedule.Lecturer.Email,
		},
	}, nil 
}