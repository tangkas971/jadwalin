package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type SubjectRepository interface {
	Create(subject *model.Subject) error
	FindByCode(code string)(*model.Subject, error)
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
		return nil, err
	}

	return &subject, nil 
}