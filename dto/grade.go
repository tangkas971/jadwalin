package dto

type GradeRequestDTO struct {
	Code    string `json:"code" binding:"required"`
	Name    string `json:"name" binding:"required"`
	ProdiId int    `json:"prodi_id" binding:"required"`
}

type GradeResponseDTO struct {
	Id      int    `json:"id" binding:"required"`
	Code    string `json:"code" binding:"required"`
	Name    string `json:"name" binding:"required"`
	ProdiId int    `json:"prodi_id" binding:"required"`
	Prodi   ProdiResponseDTO
}