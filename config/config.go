package config

import "gorm.io/gorm"

var (
    logger *Logger
	db *gorm.DB
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	//Initialize logger
	logger = NewLogger(p)
	return logger
}