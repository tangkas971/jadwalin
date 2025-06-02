package services

import (
	"jadwalin/dto"
	"jadwalin/model"
	"jadwalin/repository"
)

type ProdiService interface {
	Create(input dto.ProdiRequestDTO) error
}

type prodiService struct {
	prodiRepo repository.ProdiRepository
}

func NewProdiService(prodiRepo repository.ProdiRepository) ProdiService {
	return &prodiService{
		prodiRepo: prodiRepo,
	}
}

func (s *prodiService) Create(input dto.ProdiRequestDTO) error {
	prodi := model.Prodi{
		Code: input.Code,
		Name: input.Name,
	}
	err := s.prodiRepo.CreateProdi(&prodi)
	if err != nil {
		return err
	}

	return nil 
}