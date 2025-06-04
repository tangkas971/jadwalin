package model

import "time"

type Task struct {
	Id int
	Title string
	Description string
	Deadline time.Time

	SubjectId int
	Subject Subject

	LecturerId int
	Lecturer User 

	CreatedAt time.Time
	UpdatedAt time.Time

}