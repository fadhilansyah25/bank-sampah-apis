package LoginHandlerTest

import (
	"bytes"
	"encoding/json"
	"golang-final-project/Configs"
	"golang-final-project/Models/Response"
	"golang-final-project/Models/UserLogins"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	Configs.Connection(Configs.DBConfig{
		Host:     "localhost",
		User:     "root",
		Password: "",
		Port:     "3306",
		DBName:   "go_bank_sampah_test",
	})
	return Configs.DB
}

func Test_CreateUserLogin_OK(t *testing.T) {
	db := setupTestDB()

	db.Exec("DELETE FROM login_data_users;")
	db.Exec("ALTER TABLE login_data_users AUTO_INCREMENT = 1;")

	a := assert.New(t)

	userLogin := UserLogins.LoginDataUsers{
		UserId:   1,
		Email:    "fadhilansyah25@gmail.com",
		Username: "fadilardiansyah",
		Password: "bismillah3x",
	}

	reqBody, err := json.Marshal(userLogin)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateUserLoginRouter(db, bytes.NewBuffer(reqBody))
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

func Test_Login_OK(t *testing.T) {
	db := setupTestDB()

	a := assert.New(t)
	login := UserLogins.Login{
		Email:    "fadhilansyah25@gmail.com",
		Username: "fadilardiansyah",
		Password: "bismillah3x",
	}

	reqBody, err := json.Marshal(login)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setLoginRouter(db, bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
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
		Message: "access complete",
	}
	a.Equal(expected, actual)
}

func Test_UpdateLogin_OK(t *testing.T) {
	db := setupTestDB()

	a := assert.New(t)
	userLogin := UserLogins.LoginDataUsers{
		UserId:   1,
		Email:    "fadhilansyah25@gmail.com",
		Username: "fadhilansyah25",
		Password: "sholawat5x",
	}

	reqBody, err := json.Marshal(userLogin)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setUpdateLoginRouter(db, bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
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

func Test_GetUserLoginByID_OK(t *testing.T) {
	db := setupTestDB()

	a := assert.New(t)

	req, w := setGetUserLoginByIDRouter(db)

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
