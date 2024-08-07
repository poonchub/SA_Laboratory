package entity

import (
	"time"

	"gorm.io/gorm"
)

type EnrollmentSchedule struct {
	gorm.Model
	ScheduleDate	time.Time
	SubjectType		string
	StudyDurationHR	time.Time

	SubjectID	uint	
	TeacherID	string

}