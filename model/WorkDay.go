package model

import (
	"gorm.io/gorm"
	"time"
)

type WorkDay struct {
	gorm.Model
	Date       time.Time
	WorkShifts []WorkShift
}
