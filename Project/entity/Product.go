package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ProductName		string
	Description		string
	PricePerPiece	float32
	Stock			int

	OrderItem		[]OrderItem	`gorm:"foreignKey:ProductID"`
}