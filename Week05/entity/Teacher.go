package entity

import (
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	CreatedAt time.Time
  	UpdatedAt time.Time
  	DeletedAt gorm.DeletedAt `gorm:"index"`
	TeacherID	string	`gorm:"primaryKey"`
	FirstName	string
	LastName	string
	Email		string
	PhoneNumber	string

	EnrollmentSchedule	[]EnrollmentSchedule	`gorm:"foreignKey:TeacherID"`

}