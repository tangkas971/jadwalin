package model

import "time"

type Schedule struct {
    Id         int
    Day        string
    StartTime  string
    EndTime    string

    SubjectId  int
	Subject    Subject 

    LecturerId int
    Lecturer   User 

    GradeId int
    Grade Grade

    ProdiId int
    Prodi Prodi
    
    CreatedAt  time.Time
    UpdatedAt  time.Time
}
