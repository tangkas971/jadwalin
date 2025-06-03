package services

import (
	"fmt"
	"jadwalin/dto"
	"jadwalin/model"
	"jadwalin/repository"
	"jadwalin/utils"
)

type ProdiService interface {
	Create(userRole string, input dto.ProdiRequestDTO) error
}

type prodiService struct {
	prodiRepo repository.ProdiRepository
}

func NewProdiService(prodiRepo repository.ProdiRepository) ProdiService {
	return &prodiService{
		prodiRepo: prodiRepo,
	}
}

func (s *prodiService) Create(userRole string, input dto.ProdiRequestDTO) error {
	// Cek apakah user role adalah admin
	err := utils.RoleCheck(userRole, "admin")
	if err != nil {
		return err
	}

	// cek apakah kode prodi sudah ada 
	existingProdi, err := s.prodiRepo.FindByCode(input.Code)
	if err != nil {
		return err
	}

	if existingProdi != nil {
		return fmt.Errorf("kode prodi yang anda masukkan sudah terdaftar")
	}
	
	prodi := model.Prodi{
		Code: input.Code,
		Name: input.Name,
	}
	err = s.prodiRepo.CreateProdi(&prodi)
	if err != nil {
		return err
	}

	return nil 
}