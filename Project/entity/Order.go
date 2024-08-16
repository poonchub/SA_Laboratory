package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderDate	time.Time
	TotalPrice	*float32
	Status		string

	OrderItem	[]OrderItem	`gorm:"foreignKey:OrderID"`

	AddressID	uint
	CustomerID	uint
}