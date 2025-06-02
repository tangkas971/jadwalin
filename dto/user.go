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