package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type GradeRepository interface {
	CreateGrade(grade *model.Grade) error
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