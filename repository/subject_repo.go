package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type SubjectRepository interface {
	Create(subject *model.Subject) error
	FindByCode(code string)(*model.Subject, error)
	Update(subject *model.Subject) error
	FindById(id int)(*model.Subject, error)
	GetAll()([]*model.Subject, error)
	Delete(id int) error
}

type subjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	return &subjectRepository{
		db: db,
	}
}

func (r *subjectRepository) Create(subject *model.Subject) error {
	err := r.db.Create(subject).Error
	return err
}

func (r *subjectRepository) FindByCode(code string)(*model.Subject, error){
	var subject model.Subject
	err := r.db.Where("code = ?", code).First(&subject).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil, nil 
		}
		return nil, err
	}

	return &subject, nil 
}

func (r *subjectRepository) GetAll()([]*model.Subject, error){
	var subjects []*model.Subject

	err := r.db.Preload("Prodi").Find(&subjects).Error
	if err != nil {
		return nil, err
	}

	return subjects, nil 
}

func (r *subjectRepository) Delete(id int) error {
	err := r.db.Delete(&model.Subject{}, id).Error
	return err
}

func (r *subjectRepository) FindById(id int)(*model.Subject, error){
	var subject model.Subject
	err := r.db.Where("id = ? ", id).First(&subject).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil, nil
		}
		return nil, err
	}

	return &subject, nil 
}

func (r *subjectRepository) Update(subject *model.Subject) error{
	err := r.db.Save(subject).Error
	return err
}