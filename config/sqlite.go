package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func InitializeSQLite() (*gorm.DB, error){
	logger := GetLogger("sqlite")

	db, err := gorm.Open(sqlite.Open("./db/openings.db"), &gorm.Config{})
}