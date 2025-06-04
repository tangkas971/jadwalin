package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id int)(*model.User, error)
	FindByEmail(email string)(*model.User, error)
	FindByRole(roleUser string)([]*model.User, error)
	FindByNim(nim int)(*model.Student, error)
	FindByNip(nip int)(*model.Lecturer, error)
	FindStudentByGradeAndProdi(grade_id int, prodi_id int)([]*model.Student, error)
	FindAll()([]*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll()([]*model.User, error){
	var users []*model.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil 
}

func (r *userRepository) FindById(id int)(*model.User, error) {
	var user *model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string)(*model.User, error){
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil, nil 
		}
		return nil, err
	}

	return &user, nil
}

func(r *userRepository) FindByNim(nim int)(*model.Student, error){
	var student *model.Student
	err := r.db.Where("nim = ?", nim).First(&student).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil, nil 
		}
		return nil, err
	}
	return student, nil 
}

func (r *userRepository) FindByNip(nip int)(*model.Lecturer, error){
	var lecturer model.Lecturer
	err := r.db.Where("nip = ?", nip).First(&lecturer).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil, nil 
		}
		return nil, err
	}
	return &lecturer, nil 
}

func (r *userRepository) FindByRole(roleUser string)([]*model.User, error){
	var users []*model.User
	err := r.db.Where("role = ?", roleUser).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil 
}

func (r *userRepository) FindStudentByGradeAndProdi(grade_id int, prodi_id int)([]*model.Student, error){
	var students []*model.Student
	err := r.db.Where("prodi_id = ?", prodi_id).Where("grade_id = ?", grade_id).Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil 
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id int) error {
	return r.db.Delete(&model.User{}, id).Error
}