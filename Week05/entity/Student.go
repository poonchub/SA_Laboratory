package entity

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	CreatedAt time.Time
  	UpdatedAt time.Time
  	DeletedAt gorm.DeletedAt `gorm:"index"`
	StudentID	string	`gorm:"primaryKey"`
	FirstName	string
	LastName	string
	DateOfBirth	time.Time
	Email		string
	PhoneNumber	string
}