package Middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-final-project/Models/RequestLogging"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func Log(c echo.Context, reqBody, resBody []byte) {
	var data map[string]interface{}

	if err := json.Unmarshal(resBody, &data); err != nil {
		panic(err)
	}

	halo := data["message"].(string)

	id, _ := GetClaimsUserId(c)

	reqLogDB := RequestLogging.RequestLog{
		Time:     time.Now(),
		UserId:   id,
		Host:     c.Request().Host,
		Method:   c.Request().Method,
		Url:      c.Request().RequestURI,
		Status:   c.Response().Status,
		Message:  halo,
		RemoteIp: c.Request().RemoteAddr,
	}

	insert(&reqLogDB)

}

func connect() (*mongo.Database, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	username := os.Getenv("DB_MONGO_USERNAME")
	password := os.Getenv("DB_MONGO_PASSWORD")
	mongoUri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.u1fn4.mongodb.net/bank_sampah?retryWrites=true&w=majority", username, password)

	clientOptions := options.Client()
	clientOptions.ApplyURI(mongoUri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("bank_sampah"), nil
}

func insert(requestLog *RequestLogging.RequestLog) {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("API_log_request").InsertOne(ctx, &requestLog)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Log Saved!")
}
