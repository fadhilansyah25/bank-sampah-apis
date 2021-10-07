package main

import (
	"golang-final-project/Configs/Database"
	"golang-final-project/Routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var db Database.DBConfig

	errorEnv := godotenv.Load(".env")

	if errorEnv != nil {
		log.Fatalf("Error loading .env file")
	}

	db.Host = os.Getenv("DB_HOST")
	db.Port = os.Getenv("DB_PORT")
	db.User = os.Getenv("DB_USER")
	db.Password = os.Getenv("DB_PASSWORD")
	db.DBName = os.Getenv("DB_NAME")

	Database.Connection(db)
	e := Routes.RouteVersion1()

	port := os.Getenv("PORT")
	e.Start(":" + port)
}
