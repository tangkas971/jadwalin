package model

type Grade struct {
	Id      int
	Code    string
	Name    string
	ProdiId int
	Prodi   Prodi
}