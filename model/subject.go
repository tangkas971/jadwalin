package model

import "time"

type Subject struct {
	Id        int
	Code      string
	Name      string
	ProdiId	  int 
	Prodi	  Prodi
	CreatedAt time.Time
	UpdatedAt time.Time
}