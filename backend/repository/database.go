package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

type AppConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() AppConfig {
	return AppConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

var defaultConfig = AppConfig{
	DBHost:     "localhost",
	DBPort:     "5432",
	DBUser:     "myuser",
	DBPassword: "mypassword",
	DBName:     "mydb",
}

func ConnectToDB() (*gorm.DB, error) {
	once.Do(func() {
		cfg := LoadConfig()
		if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBPassword == "" {
			cfg = defaultConfig
		}
		dbConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
		db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
		if err != nil {
			dbConn = nil
			return
		}
		dbConn = db
	})

	if dbConn == nil {
		return nil, fmt.Errorf("failed to connect to the database")
	}

	return dbConn, nil
}
