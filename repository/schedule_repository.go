package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Create(schedule *model.Schedule) error
	GetAll()([]*model.Schedule, error)
	Delete(id int) error
	Update(schedule *model.Schedule) error 
	FindById(id int)(*model.Schedule, error)
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository{
	return &scheduleRepository{db: db}
}

func (r *scheduleRepository) Create(schedule *model.Schedule) error {
	err := r.db.Create(schedule).Error
	return err
}

func (r *scheduleRepository) GetAll()([]*model.Schedule, error){
	var schedules []*model.Schedule
	err := r.db.Preload("Subject").Preload("Prodi").Preload("Grade").Preload("Lecturer").Find(&schedules).Error
	if err != nil {
		return nil, err
	}

	return schedules, nil 
}

func (r *scheduleRepository) Delete(id int) error {
	err := r.db.Delete(&model.Schedule{}, id).Error
	return err
}

func (r *scheduleRepository) Update(schedule *model.Schedule) error{
	err := r.db.Save(schedule).Error
	return err
}

func (r *scheduleRepository) FindById(id int)(*model.Schedule, error){
	var schedule model.Schedule
	err := r.db.Where("id = ?", id).First(&schedule).Error
	if err != nil {
		return nil, err
	}

	return &schedule, nil 
}