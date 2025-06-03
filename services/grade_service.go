package services

import (
	"fmt"
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
	// Cek apakah user role adalah admin
	err := utils.RoleCheck(userRole, "admin")
	if err != nil {
		return err
	}

	// cek apakah kode grade sudah terdaftar atau belum
	existingGrade, err := s.gradeRepo.FindByCode(input.Code) 
	if err != nil {
		return err
	}

	if existingGrade != nil {
		return fmt.Errorf("code grade yang anda masukkan sudah terdaftar")
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