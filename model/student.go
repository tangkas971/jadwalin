package model

type Student struct {
	UserId int
	Nim    int

	GradeId int
	Grade  Grade

	ProdiId int
	Prodi  Prodi
	User   User
}