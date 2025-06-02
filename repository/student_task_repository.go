package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type StudentTaskRepository interface {
	Create(studentTask *model.StudentTask) error
	FindByStudentId(studentId uint)([]*model.StudentTask, error)
	Update(studentTask *model.StudentTask) error
	Delete(id int) error
}

type studentTaskRepository struct {
	db *gorm.DB
}

func NewStudentTaskRepository(db *gorm.DB) StudentTaskRepository {
	return &studentTaskRepository{
		db: db,
	}
}

func (r *studentTaskRepository) Create(studentTask *model.StudentTask) error {
	err := r.db.Save(studentTask).Error
	return err
}

func (r *studentTaskRepository) FindByStudentId(studentId uint)([]*model.StudentTask, error) {
	var studentTasks []*model.StudentTask
	err := r.db.
	Where("student_id = ?", studentId).
	Preload("Task").
	Preload("Task.Schedule").
	Preload("Task.Lecturer").
	Preload("Task.Schedule.Subject").
	Find(&studentTasks).Error

	if err != nil {
		return nil, err
	}

	return studentTasks, nil 
}

func (r *studentTaskRepository) Update(studentTask *model.StudentTask) error{
	err := r.db.Save(studentTask).Error
	return err
}

func (r *studentTaskRepository) Delete(id int) error {
	err := r.db.Delete(&model.StudentTask{}, id).Error
	return err	
}