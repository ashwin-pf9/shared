package db

import (
	"fmt"
	"log"
	"os"
	"pipeline-shared/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global DB instance

// InitDatabase initializes the database connection
func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	user := os.Getenv("DB_USER")
	database := os.Getenv("DB_DATABASE")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=require", host, user, password, database, port)

	dialector := postgres.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// making database instance Global
	DB = db

	// Auto-migrate models
	err = DB.AutoMigrate(
		&domain.Role{},
		&domain.Permission{},
		&domain.RolePermission{},
		&domain.Profile{},
		&domain.Pipeline{},
		&domain.Stage{},
		&domain.PipelineExecution{},
		&domain.StageExecution{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database connected and migrated successfully!")

	//-------------

	// Read the SQL file
	sqlFile := "/Users/ashwintirpude/DMP2S/internal/infrastructure/db/public.user.sql"
	content, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
	}

	// Get the raw database connection from GORM
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get raw DB from GORM: %v", err)
	}

	// Execute the SQL file
	_, err = sqlDB.Exec(string(content))
	if err != nil {
		log.Fatalf("Failed to execute SQL file: %v", err)
	}

	fmt.Println("SQL file executed successfully.")

}
