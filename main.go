package main

import (
	"golang-final-project/Configs"
	"golang-final-project/Routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	errorEnv := godotenv.Load(".env")

	if errorEnv != nil {
		log.Fatalf("Error loading .env file")
	}

	var dbConf Configs.DBConfig

	dbConf.Host = os.Getenv("DB_HOST")
	dbConf.Port = os.Getenv("DB_PORT")
	dbConf.User = os.Getenv("DB_USER")
	dbConf.Password = os.Getenv("DB_PASSWORD")
	dbConf.DBName = os.Getenv("DB_NAME")

	Configs.Connection(dbConf)
	e := Routes.RouteVersion1()
	e.Start(":8080")
}
