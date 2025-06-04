package dto

import "time"

type TaskRequestDTO struct{
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Deadline time.Time `json:"deadline" binding:"required"`
	SubjectId int `json:"subject_id" binding:"required"`
	LecturerId int `json:"lecturer_id"`
}

type TaskResponseDTO struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Deadline time.Time `json:"deadline"`
	
	SubjectId int `json:"subject_id"`
	Subject SubjectResponseDTO

	LecturerId int `json:"lecturer_id"`
	Lecturer LecturerResponseDTO
}