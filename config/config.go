package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = initializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing database %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {

	logger = NewLogger(p)
	return logger
}
