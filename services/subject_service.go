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
	GetAll()([]dto.SubjectResponseDTO, error)
	Delete(id int) error
	Update(id int, input dto.SubjectRequestDTO) error
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

func (s *subjectService) GetAll()([]dto.SubjectResponseDTO, error){
	subjects, err := s.subjectRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var subjectDTOs []dto.SubjectResponseDTO
	for _, subject := range subjects{
		subjectDTO := dto.SubjectResponseDTO{
			Id: subject.Id,
			Code: subject.Code,
			Name: subject.Name,
			ProdiId: subject.ProdiId,
			Prodi: dto.ProdiResponseDTO{
				Name: subject.Prodi.Name,
			},
			CreateAt: subject.CreatedAt,
			UpdatedAt: subject.UpdatedAt,
		}
		subjectDTOs = append(subjectDTOs, subjectDTO)
	}

	return subjectDTOs, nil 
}

func (s *subjectService) Delete(id int) error{
	err := s.subjectRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil 
}

func (s *subjectService) Update(id int, input dto.SubjectRequestDTO) error{
	existingSubject, err := s.subjectRepo.FindById(id)
	if err != nil {
		return err
	}

	existingSubject.Code = input.Code
	existingSubject.Name = input.Name
	existingSubject.ProdiId = input.ProdiId

	err = s.subjectRepo.Update(existingSubject)
	if err != nil {
		return err
	}

	return nil
}