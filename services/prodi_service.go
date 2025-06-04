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
	GetAll()([]dto.ProdiResponseDTO, error)
	Update(userRole string, id int, input dto.ProdiRequestDTO) error
	Delete(userRole string, id int) error
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

func (s *prodiService) GetAll()([]dto.ProdiResponseDTO, error){
	// ambil data prodis
	prodis, err := s.prodiRepo.GetAll()
	if err != nil {
		return []dto.ProdiResponseDTO{}, err
	}

	// masukkan ke dto 
	var prodiDTOs []dto.ProdiResponseDTO
	for _, prodi := range prodis{
		prodiDTO := dto.ProdiResponseDTO{
			Id: prodi.Id,
			Code: prodi.Code,
			Name: prodi.Name,
		}
		prodiDTOs = append(prodiDTOs, prodiDTO)
	}

	return prodiDTOs, nil 
}

func (s *prodiService) Update(userRole string, id int, input dto.ProdiRequestDTO) error{
	// cek role user
	err := utils.RoleCheck(userRole, "admin")
	if err != nil {
		return err
	}

	existingProdi, err := s.prodiRepo.FindyById(id)
	if err != nil {
		return err
	}

	existingProdi.Code = input.Code
	existingProdi.Name = input.Name

	err = s.prodiRepo.Update(existingProdi)
	if err != nil {
		return err
	}

	return nil
}

func (s *prodiService) Delete(userRole string, id int) error{
	// cek role user
	err := utils.RoleCheck(userRole, "admin")
	if err != nil {
		return err
	}
	
	err = s.prodiRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil 
}