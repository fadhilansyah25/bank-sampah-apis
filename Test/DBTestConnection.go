package Test

import (
	"golang-final-project/Configs/Database"
	"os"

	"gorm.io/gorm"
)

func SetUpTestDB() *gorm.DB {
	// errorEnv := godotenv.Load(".env")

	// if errorEnv != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	Database.Connection(Database.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   "go_bank_sampah_test",
	})
	return Database.DB
}
