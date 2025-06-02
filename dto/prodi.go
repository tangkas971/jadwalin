package dto

type ProdiRequestDTO struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type ProdiResponseDTO struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}