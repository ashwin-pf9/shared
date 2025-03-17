package migrations

import (
	"log"
	"os"

	"github.com/ashwin-pf9/shared/domain"
	"gorm.io/gorm"
)

// RunAutoMigrations handles GORM-based migrations.
func RunAutoMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
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
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("AutoMigrations completed successfully.")
}

// RunSQLFileMigration executes SQL from a file.
func RunSQLFileMigration(db *gorm.DB, filepath string) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get raw DB from GORM: %v", err)
	}

	_, err = sqlDB.Exec(string(content))
	if err != nil {
		log.Fatalf("Failed to execute SQL file: %v", err)
	}

	log.Println("SQL file migration executed successfully.")
}
