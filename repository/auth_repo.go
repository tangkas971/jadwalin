package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateStudent(student *model.Student) error
	CreateLecturer(lecturer *model.Lecturer) error
	CreateUser(user *model.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) CreateStudent(student *model.Student) error {
	err := r.db.Create(student).Error
	return err
}

func (r *authRepository) CreateLecturer(lecturer *model.Lecturer) error{
	err := r.db.Create(lecturer).Error
	return err
}

func (r *authRepository) CreateUser(user *model.User) error {
	err := r.db.Save(user).Error
	return err
}

