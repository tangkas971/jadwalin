package model

import "time"

type StudentTask struct {
	Id          int
	StudentId   int
	Student     User 

	TaskId      int
	Task        Task 

	Status      string
	
	SubmittedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
