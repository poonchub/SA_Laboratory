package entity

import (
	"time"

	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	SubjectName	string
	SubjectType	string
	TotalStudyHours	time.Time
	Note		string

	EnrollmentSchedule	[]EnrollmentSchedule	`gorm:"foreignKey:SubjectID"`

	CourseTypeID	uint
}