package model

import (
	"gorm.io/gorm"
	"time"
)

type WorkShift struct {
	gorm.Model
	EmployeeID      uint
	WorkShiftTypeID uint
	WorkDayID       uint
	StartTime       time.Time
	EndTime         time.Time
	Duration        time.Duration
}
