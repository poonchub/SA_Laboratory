package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type students struct {
  gorm.Model
  StudentID	string
  Name string
  Team string
}

func main() {
  db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&students{})

  // Create
  db.Create(&students{StudentID: "B6525163", Name: "Poonchub Nanawan", Team: "06"})

  // Read
//   var product students
//   db.First(&product, 1) // find product with integer primary key
//   db.First(&product, "code = ?", "B6525163") // find product with code D42

  // Update - update product's price to 200
  // db.Model(&product).Update("Price", 200)
  // Update - update multiple fields
  // db.Model(&product).Updates(students{Price: 200, Code: "F42"}) // non-zero fields
  // db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  // db.Delete(&product, 1)
}