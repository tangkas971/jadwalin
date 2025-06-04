package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type GradeRepository interface {
	CreateGrade(grade *model.Grade) error
	GetAll()([]*model.Grade, error)
	FindByCode(code string)(*model.Grade, error)
	FindById(id int)(*model.Grade, error)
	Update(grade *model.Grade) error
	Delete(id int) error
}

type gradeRepository struct {
	db *gorm.DB
}

func NewGradeRepository(db *gorm.DB) GradeRepository {
	return &gradeRepository{
		db: db,
	}
}

func (r *gradeRepository) CreateGrade(grade *model.Grade) error {
	err := r.db.Create(grade).Error
	return err
}

func (r *gradeRepository) GetAll()([]*model.Grade, error){
	var grades []*model.Grade
	err := r.db.Preload("Prodi").Find(&grades).Error
	if err != nil {
		return nil, err
	}

	return grades, nil 
}

func (r *gradeRepository) FindByCode(code string)(*model.Grade, error){
	var grade model.Grade
	err := r.db.Where("code = ?", code).First(&grade).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil, nil 
		}
		return nil, err
	}
	return &grade, nil 
}

func (r *gradeRepository) FindById(id int)(*model.Grade, error){
	var grade model.Grade

	err := r.db.Where("id = ?", id).First(&grade).Error
	if err != nil {
		return nil, err
	}

	return &grade, nil 
}

func (r *gradeRepository) Delete(id int) error{
	err := r.db.Delete(&model.Grade{}, id).Error
	return err 
} 

func (r *gradeRepository) Update(grade *model.Grade) error{
	err := r.db.Save(grade).Error
	return err
}