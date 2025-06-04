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
	GetAll()([]dto.GradeResponseDTO, error)
	Update(userRole string, id int, input dto.GradeRequestDTO) error
	Delete(userRole string, id int) error
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

func (s *gradeService) GetAll()([]dto.GradeResponseDTO, error){
	grades, err := s.gradeRepo.GetAll()
	if err != nil {
		return []dto.GradeResponseDTO{}, err
	}

	var gradeDTOs []dto.GradeResponseDTO

	for _, grade := range grades{
		gradeDTO := dto.GradeResponseDTO{
			Id: grade.Id,
			Code: grade.Code,
			Name: grade.Name,
			ProdiId: grade.Prodi.Id,
			Prodi: dto.ProdiResponseDTO{
				Name: grade.Prodi.Name,
			},
		}
		gradeDTOs = append(gradeDTOs, gradeDTO)
	}

	return gradeDTOs, nil 
}

func (s *gradeService) Delete(userRole string, id int) error{
	err := utils.RoleCheck(userRole, "admin")
	if err != nil {
		return err
	}
	err = s.gradeRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil 
}

func (s *gradeService) Update(userRole string, id int, input dto.GradeRequestDTO) error{
	
	err := utils.RoleCheck(userRole, "admin")
	if err != nil {
		return err
	}
	existingGrade, err := s.gradeRepo.FindById(id)
	if err != nil{
		return err
	}

	existingGrade.Code = input.Code
	existingGrade.Name = input.Name
	existingGrade.ProdiId = input.ProdiId

	err = s.gradeRepo.Update(existingGrade)
	if err != nil {
		return err
	}

	return nil
}