package models

import (
	"fmt"
	"log"

	"BoardGame/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configs.Cfg.DB_HOST,
		configs.Cfg.DB_USER,
		configs.Cfg.DB_PASSWORD,
		configs.Cfg.DB_NAME,
		configs.Cfg.DB_PORT,
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	log.Println("Database connected")

	DB = database
}
