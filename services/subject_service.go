package services

import (
	"fmt"
	"jadwalin/dto"
	"jadwalin/model"
	"jadwalin/repository"
	"jadwalin/utils"
)

type SubjectService interface {
	Create(userRole string, input dto.SubjectRequestDTO) error
}

type subjectService struct {
	subjectRepo repository.SubjectRepository
}

func NewSubjectService(subjectRepo repository.SubjectRepository) SubjectService{
	return &subjectService{
		subjectRepo: subjectRepo,
	}
}

func (s *subjectService) Create(userRole string, input dto.SubjectRequestDTO) error {
	// cek apakah user role adalah admin
	err := utils.RoleCheck(userRole, "admin")
	if err != nil {
		return fmt.Errorf("role anda %s! Hanya admin yang dapat menambahkan subject", userRole)
	}

	// cek kode subject sudah terdaftar atau belum
	existingSubject, err := s.subjectRepo.FindByCode(input.Code)
	if err != nil {
		return err
	}

	if existingSubject != nil {
		return fmt.Errorf("code subject yang kamu masukkan sudah terdaftar")
	}

	subject := model.Subject{
		Code: input.Code,
		Name: input.Name,
		ProdiId: input.ProdiId,
	}

	err = s.subjectRepo.Create(&subject)
	if err != nil {
		return err
	}

	return nil 
}