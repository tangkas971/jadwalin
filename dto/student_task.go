package dto

import "time"

type StudentTaskResponse struct {
	Id          int 			`json:"id"`
	TaskId      int 			`json:"task_id"`
	Task        TaskResponseDTO
	Status      string 			`json:"status"`
	SubmittedAt *time.Time 		`json:"submitted_at"`
}