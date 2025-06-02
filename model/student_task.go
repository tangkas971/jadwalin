package model

import "time"

type StudentTask struct {
	Id          int
	StudentId   int
	Student     User `gorm:"foreignKey:StudentId;references:Id"`
	TaskId      int
	Task        Task `gorm:"foreignKey:TaskId;references:Id"`
	Status      string
	SubmittedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
