package utils

import (
	"fmt"

	"github.com/pejuang-awan/BE-Tournament-Manager/config"
	"github.com/pejuang-awan/BE-Tournament-Manager/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitializeDatabase() {

	fmt.Println("Database initialized")

	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
			config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBDatabase,
			config.AppConfig.DBUsername, config.AppConfig.DBPassword, config.AppConfig.SSLMode),
	), &gorm.Config{})

	if err != nil {
		panic("[ERROR] Failed to initialize database")
	}

	err = db.AutoMigrate(
		&models.Tournament{},
		&models.Game{},
	)

	if err != nil {
		fmt.Println("[ERROR] Failed to migrate database")
	}

	fmt.Println("Database connected")

	Database = db
}
