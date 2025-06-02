package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	FindAll()([]*model.Task, error)
	FindById(id int)(*model.Task, error)
	Create(task *model.Task) error
	Update(task *model.Task) error
	Delete(id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) FindAll()([]*model.Task, error){
	var tasks []*model.Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil 
}

func (r *taskRepository) FindById(id int)(*model.Task, error){
	var task *model.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return task, nil 
}

func (r *taskRepository) Create(task *model.Task) error {
	err := r.db.Save(task).Error
	return err
}

func (r *taskRepository) Update(task *model.Task) error {
	err := r.db.Save(task).Error
	return err
}

func (r *taskRepository) Delete(id int) error {
	err := r.db.Delete(&model.Task{}, id).Error
	return err
}