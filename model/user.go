package model

import "time"

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	Role 	  string
	CreatedAt time.Time
	UpdatedAt time.Time
}