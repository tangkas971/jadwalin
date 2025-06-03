package dto

type AdminRegisterRequest struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Role            string `json:"role" binding:"required,oneof=mahasiswa dosen admin"`
}

type StudentRegisterRequest struct {
	Name            string `json:"name" binding:"required"`
	NIM             int    `json:"nim" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Role            string `json:"role" binding:"required,oneof=mahasiswa dosen admin"`
	GradeId         int `json:"grade_id" binding:"required"`
	ProdiId         int `json:"prodi_id" binding:"required"`
}

type LecturerRegisterRequest struct {
	Name            string `json:"name" binding:"required"`
	NIP             int    `json:"nip" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Role            string `json:"role" binding:"required,oneof=mahasiswa dosen admin"`
	ProdiId	        int	   `json:"prodi_id" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDTO struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}