package dto

import "time"

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	LecturerId  int `json:"lecturer_id" binding:"required"`
	ScheduleId  int `json:"schedule_id" binding:"required"`
}

type TaskResponseDTO struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	LecturerId  int                 `json:"lecturer_id"`
	Lecturer    UserResponseDTO     `json:"User"`
	ScheduleId  int                 `json:"schedule_id"`
	Schedule    ScheduleResponseDTO `json:"schedule"`
}

type UpdateTaskRequestDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}