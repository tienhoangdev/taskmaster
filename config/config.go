package config

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"golang_project_base/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func InitDB() {

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		host, username, database, password, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB = db

	// Automatically migrate your schema
	if err := db.AutoMigrate(&models.Habit{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
