package services

import (
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
	err := utils.RoleCheck(userRole, "admin")
	if err != nil {
		return err
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