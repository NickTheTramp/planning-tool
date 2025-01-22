package model

import (
	"gorm.io/gorm"
	"time"
)

type WorkShiftType struct {
	gorm.Model
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
}
