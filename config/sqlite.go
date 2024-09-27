package config

import (
	"github.com/Calebbuffleben/api-with-golang/schemas"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func InitializeSQLite() (*gorm.DB, error){
	logger := GetLogger("sqlite")

	db, err := gorm.Open(sqlite.Open("./db/openings.db"), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite auto migration: %v", err)
		return nil, err
	}

	return db, nil
}