package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Create(schedule *model.Schedule) error
	FindAll()([]*model.Schedule, error)
	FindById(id int)(*model.Schedule, error)
	Delete(id int) error
	Update(schedule *model.Schedule) error
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository{
	return &scheduleRepository{db: db}
}

func (r *scheduleRepository) Create(schedule *model.Schedule) error {
	err := r.db.Save(schedule).Error
	return err
}

func (r *scheduleRepository) Delete(id int) error{
	return r.db.Delete(&model.Schedule{}, id).Error
}

func (r *scheduleRepository) FindAll()([]*model.Schedule, error){
	var schedules []*model.Schedule
	err := r.db.Preload("Lecturer").Preload("Subject").Find(&schedules).Error
	if err != nil {
		return nil, err
	}

	return schedules, nil 
}

func (r *scheduleRepository) FindById(id int)(*model.Schedule, error){
	var schedule *model.Schedule

	err := r.db.First(&schedule, id).Error
	if err != nil{
		return nil, err
	}

	return schedule, nil
}

func (r *scheduleRepository) Update(schedule *model.Schedule) error {
	err := r.db.Save(&schedule).Error
	return err
}