package dto

type CreateScheduleRequestDTO struct {
	Day        string `json:"day" binding:"required"`
	StartTime  string `json:"start_time" binding:"required"`
	EndTime    string `json:"end_time" binding:"required"`
	SubjectId  int    `json:"subject_id" binding:"required"`
	LecturerId int    `json:"lecturer_id" binding:"required"`
}

type UpdateScheduleRequestDTO struct {
	Day       string `json:"day" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
}

type ScheduleResponseDTO struct {
	Id         int                 `json:"id"`
	Day        string              `json:"day"`
	StartTime  string              `json:"start_time"`
	EndTime    string              `json:"end_time"`
	SubjectId  int                 `json:"subject_id"`
	Subject    SubjectResponseDTO  `json:"subject"`
	LecturerId int                 `json:"lecturer_id"`
	Lecturer   LecturerResponseDTO `json:"lecturer"`
}
