package entity

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Prefix		string
	FirstName	string
	LastName	string
	Email		string
	Password	string
	Birthday	time.Time

	Address		[]Address	`gorm:"foreignKey:CustomerID"`

	Order		[]Order	`gorm:"foreignKey:CustomerID"`
}