package model

import "time"

type StudentSchedules struct {
	Id int

	StudentId int
	Student   User

	ScheduleId int
	Schedule   Schedule

	CreatedAt time.Time
	UpdatedAt time.Time
}