package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	// return errors.New("fake Error")
	return nil
}

func GetLogger(p string) *Logger {
	// init logger
	logger := NewLogger(p)
	return logger
}
