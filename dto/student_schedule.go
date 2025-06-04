package dto

type StudentScheduleResponseDTO struct {
	Id        int `json:"id"`
	StudentId int `json:"student_id"`
	Student   StudentResponseDTO

	ScheduleId int `json:"schedule_id"`
	Schedule   ScheduleResponseDTO
}