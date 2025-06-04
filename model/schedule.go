package model

import "time"

type Schedule struct {
    Id         int
    Day        string
    StartTime  string
    EndTime    string

    SubjectId  int
	Subject    Subject `gorm:"foreignKey:SubjectId;references:Id"`

    LecturerId int
    Lecturer   User `gorm:"foreignKey:LecturerId;references:Id"`

    GradeId int
    Grade Grade

    ProdiId int
    Prodi Prodi
    
    CreatedAt  time.Time
    UpdatedAt  time.Time
}
