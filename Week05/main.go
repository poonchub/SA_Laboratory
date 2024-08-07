package main

import ("github.com/sa67-lab5/entity"
		"gorm.io/gorm"
  		"gorm.io/driver/sqlite"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&entity.Student{}, &entity.Teacher{}, &entity.Subject{}, &entity.EnrollmentSchedule{}, &entity.CourseType{})
}