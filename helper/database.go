package helper

import (
	"errors"
	"fmt"
	m "github.com/NickTheTramp/planning-tool/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() error {
	var err error

	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		return errors.New("failed to connect database")
	}

	migrateModels()

	return nil
}

func migrateModels() {
	models := []interface{}{
		&m.Employee{},
		&m.WorkDay{},
		&m.WorkShift{},
		&m.WorkShiftType{},
	}

	for _, model := range models {
		err := DB.AutoMigrate(model)

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
