package config

import (
	"fmt"
	"log"

	"gorm-basic/utils"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDatabase() (*gorm.DB, error) {
	dbHost := utils.ConfigEnv("DB_HOST")
	dbPort := utils.ConfigEnv("DB_PORT")
	dbUser := utils.ConfigEnv("DB_USER")
	dbPassword := utils.ConfigEnv("DB_PASSWORD")
	dbName := utils.ConfigEnv("DB_NAME")
	sslMode := utils.ConfigEnv("SSL_MODE")

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})	
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	return db, nil
}