package db

import (
	"fmt"
	"log"
	"os"

	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDatabase returns *gorm.DB without running migrations.
func InitDatabase() *gorm.DB {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	user := os.Getenv("DB_USER")
	database := os.Getenv("DB_DATABASE")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", host, user, password, database, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{PrepareStmt: false})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established.")
	return db
}
