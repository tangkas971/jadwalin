package dto

type CreateSubjectRequest struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type SubjectResponseDTO struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type UpdateSubjectRequest struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}