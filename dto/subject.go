package dto

import "time"

type SubjectRequestDTO struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
	ProdiId int `json:"prodi_id" binding:"required"`
}

type SubjectResponseDTO struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	ProdiId int `json:"prodi_id"`
	Prodi ProdiResponseDTO
	CreateAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateSubjectRequest struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}