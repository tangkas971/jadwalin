package dto

import "time"

type ScheduleRequestDTO struct {
	UserId     int    `json:"user_id"`
	Day        string `json:"day" binding:"required"`
	StartTime  string `json:"start_time" binding:"required"`
	EndTime    string `json:"end_time" binding:"required"`
	SubjectId  int    `json:"subject_id" binding:"required"`
	GradeId    int    `json:"grade_id" binding:"required"`
	ProdiId    int 	  `json:"prodi_id"`
	UserRole   string `json:"user_role"`
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
	ProdiId    int 				   `json:"prodi_id"`
	Prodi      ProdiResponseDTO
	GradeId    int 	 			   `json:"grade_id"`
	Grade      GradeResponseDTO
	LecturerId int                 `json:"lecturer_id"`
	Lecturer   LecturerResponseDTO `json:"lecturer"`
	CreatedAt  time.Time 		   `json:"created_at"`
	UpdateAt   time.Time 		   `json:"updated_at"`
}
