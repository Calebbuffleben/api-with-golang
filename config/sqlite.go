package config

import (
	"os"
	"github.com/Calebbuffleben/api-with-golang/schemas"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func InitializeSQLite() (*gorm.DB, error){
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err){
		logger.Info("database file not found, creating...")

		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite auto migration: %v", err)
		return nil, err
	}

	//Return de DB
	return db, nil
}