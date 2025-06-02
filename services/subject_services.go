package services

import (
	"fmt"
	"jadwalin/dto"
	"jadwalin/model"
	"jadwalin/repository"

	"gorm.io/gorm"
)

type SubjectServices interface {
	Create(roleUser string, input dto.CreateSubjectRequest)(dto.SubjectResponseDTO, error)
	FindAll()([]dto.SubjectResponseDTO, error)
	FindById(id int)(dto.SubjectResponseDTO, error)
	Delete(roleUser string, id int) error
	Update(roleUser string, id int, input dto.CreateSubjectRequest)(dto.SubjectResponseDTO, error )
}

type subjectServices struct {
	repo repository.SubjectRepository
}

func NewSubjectServices(repo repository.SubjectRepository) SubjectServices{
	return &subjectServices{
		repo: repo,
	}
}

func (s *subjectServices) Create(roleUser string, input dto.CreateSubjectRequest)(dto.SubjectResponseDTO, error){
	if roleUser != "admin" {
		return dto.SubjectResponseDTO{}, fmt.Errorf("hanya admin yang dapat menambahkan subject baru")
	}
	subject := model.Subject{
		Code: input.Code,
		Name: input.Name,
	}
	err := s.repo.Create(&subject)
	if err != nil {
		return dto.SubjectResponseDTO{}, err
	}

	return dto.SubjectResponseDTO{
		Id: subject.Id,
		Code: subject.Code,
		Name: subject.Name,
	}, nil 
}

func (s *subjectServices) FindAll()([]dto.SubjectResponseDTO, error) {
	subjects, err := s.repo.FindAll()
	if err != nil {
		return []dto.SubjectResponseDTO{}, err
	}

	var subjectDTOs []dto.SubjectResponseDTO
	for _,subject := range subjects {
		subjectDTO := dto.SubjectResponseDTO{
			Id: subject.Id,
			Code: subject.Code,
			Name: subject.Name,
		}
		subjectDTOs = append(subjectDTOs, subjectDTO)
	}

	return subjectDTOs, nil 
}

func (s *subjectServices) FindById(id int)(dto.SubjectResponseDTO, error){
	subject, err := s.repo.FindById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return dto.SubjectResponseDTO{}, fmt.Errorf("subject dengan id %d tidak ditemukan", id)
		}
		return dto.SubjectResponseDTO{}, err
	}

	return dto.SubjectResponseDTO{
		Id: subject.Id,
		Code: subject.Code,
		Name: subject.Name,
	}, nil 
}

func (s *subjectServices) Delete(roleUser string, id int) error {
	if roleUser != "admin" {
		return fmt.Errorf("hanya admin yang dapat menghapus subject")
	}
	err := s.repo.Delete(id)
	return err 
}

func (s *subjectServices) Update(roleUser string, id int, input dto.CreateSubjectRequest)(dto.SubjectResponseDTO, error){
	if roleUser != "admin" {
		return dto.SubjectResponseDTO{}, fmt.Errorf("hanya admin yang dapat mengubah subject")
	}
	
	existingSubject, err := s.repo.FindById(id)
	if err != nil {
		return dto.SubjectResponseDTO{}, err
	}

	existingSubject.Code = input.Code
	existingSubject.Name = input.Name

	err = s.repo.Update(existingSubject)
	if err != nil {
		return dto.SubjectResponseDTO{}, err
	}

	return dto.SubjectResponseDTO{
		Id: existingSubject.Id,
		Code: existingSubject.Code,
		Name: existingSubject.Name,
	}, nil 
}