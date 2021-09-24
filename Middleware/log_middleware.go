package Middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-final-project/Models/RequestLogging"
	"log"
	"time"

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
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
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
