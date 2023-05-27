package config

import (
	"os"

	"github.com/francisko-rezende/gopportunities/schemas"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB, error) {
	// dbPath := "./db/main.db"

	// check if db exists
	// _, err := os.Stat(dbPath)
	// if os.IsNotExist(err) {
	// 	logger.Info("database file not found, creating...")
	// 	// create db file and dir
	// 	err = os.MkdirAll("./db", os.ModePerm)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	file, err := os.Create(dbPath)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	file.Close()
	// }

	// Create DB and connect
	logger := GetLogger("postgres")
	err := godotenv.Load(".env")

	if err != nil {
		logger.ErrF("error loading .env file: %v", err)
	}

	dsn := "host=" + os.Getenv("POSTGRES_HOST") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=" + os.Getenv("POSTGRES_PORT") + " sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrF("postgres opening error: %v", err)
		return nil, err
	}

	// migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrF("postgres automigration error: %v", err)
		return nil, err
	}

	// return the db
	return db, nil
}
