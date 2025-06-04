package dto

import "time"

type UserResponseDTO struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	NIM       *int       `json:"nim"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LecturerResponseDTO struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type StudentResponseDTO struct {
	UserId int `json:"user_id"`
	Nim int `json:"nim"`
	GradeId int `json:"grade_id"`
	ProdiId int `json:"prodi_id"`
}