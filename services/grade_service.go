package services

import (
	"jadwalin/dto"
	"jadwalin/model"
	"jadwalin/repository"
	"jadwalin/utils"
)

type GradeService interface {
	Create(userRole string, input dto.GradeRequestDTO) error
}

type gradeService struct {
	gradeRepo repository.GradeRepository
}

func NewGradeService(gradeRepo repository.GradeRepository) GradeService {
	return &gradeService{
		gradeRepo: gradeRepo,
	}
}

func (s *gradeService) Create(userRole string, input dto.GradeRequestDTO) error{
	err := utils.RoleCheck(userRole, "admin")
	if err != nil {
		return err
	}
	grade := model.Grade{
		Code: input.Code,
		Name: input.Name,
		ProdiId: input.ProdiId,
	}
	err = s.gradeRepo.CreateGrade(&grade)
	if err != nil {
		return err
	}

	return nil 
}