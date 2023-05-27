package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitConfig() error {
	// return errors.New("fake Error")
	var err error
	// init postgres
	db, err = InitPostgres()

	if err != nil {
		return fmt.Errorf("error initializing postgres: %v", err)
	}

	return nil
}

func GetterPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// init logger
	logger := NewLogger(p)
	return logger
}
