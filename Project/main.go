package main

import (
	"Project/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("ITShop.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.Address{}, &entity.Customer{}, &entity.Order{}, entity.OrderItem{}, &entity.Product{})
}