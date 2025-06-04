package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type StudentScheduleRepository interface {
	Create(studentSchedule *model.StudentSchedules) error
}

type studentScheduleRepository struct {
	db *gorm.DB
}

func NewStudentScheduleRepository(db *gorm.DB) StudentScheduleRepository{
	return &studentScheduleRepository{
		db: db,
	}
}

func (r *studentScheduleRepository) Create(studentSchedule *model.StudentSchedules) error{
	err := r.db.Create(studentSchedule).Error
	return err
}