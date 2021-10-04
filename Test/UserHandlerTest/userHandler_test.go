package UserHandlerTest

import (
	"bytes"
	"encoding/json"
	"golang-final-project/Configs"
	"golang-final-project/Models/Response"
	"golang-final-project/Models/Users"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupDB() *gorm.DB {
	Configs.Connection(Configs.DBConfig{
		Host:     "localhost",
		User:     "root",
		Password: "",
		Port:     "3306",
		DBName:   "go_bank_sampah_test",
	})
	return Configs.DB
}

func Test_GetUsers_EmptyResult(t *testing.T) {
	db := setupDB()
	// db.Exec("DELETE FROM users")
	// db.Exec("ALTER TABLE users AUTO_INCREMENT = 1;")

	req, w := setGetUsersRouter(db)

	a := assert.New(t)
	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := Response.BaseResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := Response.BaseResponse{
		Data: []interface{}([]interface{}{}),
	}

	a.Equal(expected.Data, actual.Data)
}

func Test_CreateUser_OK(t *testing.T) {
	db := setupDB()

	db.Exec("DELETE FROM users")
	db.Exec("ALTER TABLE users AUTO_INCREMENT = 1;")

	a := assert.New(t)
	user := Users.User{
		NamaDepan:      "Mohammad",
		NamaBelakang:   "Salah",
		NIK:            "37732225059789892",
		TanggalLahir:   "1997-05-25",
		NoTelepon:      "089894983056",
		Alamat:         "Jl. Scouse Makmur No. 11, Kp. Liverpool",
		Kabupaten_Kota: "Liverpool City",
		Provinsi:       "Red Liverpool",
	}

	reqBody, err := json.Marshal(user)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateUserRouter(db, bytes.NewBuffer(reqBody))
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
		Code:    201,
		Message: "successful create data",
	}
	a.Equal(expected, actual)
}

func Test_UpdateUser_OK(t *testing.T) {
	db := setupDB()

	a := assert.New(t)
	user := Users.User{
		NamaDepan:      "Muhammad",
		NamaBelakang:   "Fadil",
		NIK:            "37732225059789892",
		TanggalLahir:   "1997-05-25",
		NoTelepon:      "089894983056",
		Alamat:         "Jl. Scouse Makmur No. 11, Kp. Liverpool",
		Kabupaten_Kota: "Liverpool City",
		Provinsi:       "Red Liverpool",
	}

	reqBody, err := json.Marshal(user)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setUpdateUserRouter(db, bytes.NewBuffer(reqBody))
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
	db := setupDB()

	a := assert.New(t)

	req, w, err := setDeleteUserRouter(db)
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
}
