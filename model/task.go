package model

import "time"

type Task struct {
	Id int
	Title string
	Description string
	Deadline time.Time

	LecturerId int
	Lecturer User  `gorm:"foreignKey:LecturerId;references:Id"`

	ScheduleId int
	Schedule Schedule  `gorm:"foreignKey:ScheduleId;references:Id"`

	CreatedAt time.Time
	UpdatedAt time.Time

}