package config

import (
	"gorm.io/gorm"
	"fmt"
)
var (
    logger *Logger
	db     *gorm.DB
)

func Init() error {
	var err error

	db, err = InitializeSQLite()

	if  err != nil {
		return fmt.Errorf("error initializing sqlite %w", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//Initialize logger
	logger = NewLogger(p)
	return logger
}