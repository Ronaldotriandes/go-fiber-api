package database

import (
	"fmt"
	"log"

	"github.com/Ronaldotriandes/go-fiber-api/src/config"
	"github.com/Ronaldotriandes/go-fiber-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	log.Println("Successfully connected to database")
	DB = db

	db.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return DB
}
