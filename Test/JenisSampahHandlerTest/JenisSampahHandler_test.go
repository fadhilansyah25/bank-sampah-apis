package JenisSampahTest

import (
	"bytes"
	"encoding/json"
	"golang-final-project/Configs/Database"
	"golang-final-project/Models/Response"
	"golang-final-project/Models/Transaction"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	errorEnv := godotenv.Load(".env")

	if errorEnv != nil {
		log.Fatalf("Error loading .env file")
	}

	Database.Connection(Database.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   "go_bank_sampah_test",
	})
	return Database.DB
}

func Test_CreateJenisSampah_OK(t *testing.T) {
	db := setupTestDB()

	db.Exec("DELETE FROM jenis_sampahs WHERE id=(SELECT MAX(id) FROM jenis_sampahs)")
	db.Exec("ALTER TABLE jenis_sampahs AUTO_INCREMENT = 1;")

	a := assert.New(t)
	jenisSampah := Transaction.JenisSampah{
		NamaJenis:  "Kain",
		MinimalQty: 1,
		Satuan:     "Kilogram",
		HargaJual:  25000,
	}

	reqBody, err := json.Marshal(jenisSampah)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateJenisSampahRouter(db, bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusCreated, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	response := Response.BaseResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		a.Error(err)
	}

	actual := Response.BaseResponse{
		Code:    response.Code,
		Message: response.Message,
	}

	expected := Response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "successful create data",
	}
	a.Equal(expected, actual)
}

func Test_GetAllJenisSampah_OK(t *testing.T) {
	db := setupTestDB()

	req, w := setGetAllJenisSampahRouter(db)

	a := assert.New(t)
	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	response := Response.BaseResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		a.Error(err)
	}

	actual := Response.BaseResponse{
		Code:    response.Code,
		Message: response.Message,
	}

	expected := Response.BaseResponse{
		Code:    http.StatusOK,
		Message: "successful retrieve data",
	}
	a.Equal(expected, actual)
}

func Test_UpdateUser_OK(t *testing.T) {
	db := setupTestDB()

	a := assert.New(t)
	jenisSampah := Transaction.JenisSampah{
		NamaJenis:  "Kain",
		MinimalQty: 1,
		Satuan:     "Kilogram",
		HargaJual:  26000,
	}

	reqBody, err := json.Marshal(jenisSampah)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setUpdateJenisSampahRouter(db, bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPut, req.Method, "HTTP request method error")
	a.Equal(http.StatusAccepted, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	response := Response.BaseResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		a.Error(err)
	}

	actual := Response.BaseResponse{
		Code:    response.Code,
		Message: response.Message,
	}

	expected := Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "successful update data",
	}
	a.Equal(expected, actual)
}

func Test_DeleteUser_OK(t *testing.T) {
	db := setupTestDB()

	a := assert.New(t)

	req, w, err := setDeleteJenisSampahRouter(db)
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodDelete, req.Method, "HTTP request method error")
	a.Equal(http.StatusAccepted, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	response := Response.BaseResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		a.Error(err)
	}

	actual := Response.BaseResponse{
		Code:    response.Code,
		Message: response.Message,
	}

	expected := Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "successful delete data",
	}
	a.Equal(expected, actual)

	db.Exec("ALTER TABLE users AUTO_INCREMENT = 1;")
}
