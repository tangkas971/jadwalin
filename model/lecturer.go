package model

type Lecturer struct {
	UserID int
	Nip    int
	ProdiId  int
	Prodi Prodi
	User   User
}