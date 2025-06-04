package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type LecturerRepository interface {
	FindById(userId int) (*model.Lecturer, error)
}

type lecturerRepository struct {
	db *gorm.DB
}

func NewLecturerRepository(db *gorm.DB) LecturerRepository {
	return &lecturerRepository{
		db: db,
	}
}

func (r *lecturerRepository)FindById(userId int) (*model.Lecturer, error){
	var lecturer model.Lecturer
	err := r.db.Where("user_id = ?", userId).First(&lecturer).Error
	if err != nil {
		return nil, err
	}

	return &lecturer, nil 
}