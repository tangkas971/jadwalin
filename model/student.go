package model

type Student struct {
	UserID int
	Nim    int

	GradeId int
	Grade  Grade

	ProdiId int
	Prodi  Prodi
	User   User
}