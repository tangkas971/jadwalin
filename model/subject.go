package model

import "time"

type Subject struct {
	Id        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}