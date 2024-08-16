package entity

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Quantity	int

	OrderID		uint
	ProductID	string
}