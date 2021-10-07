package UserHandlerTest

import (
	"bytes"
	"encoding/json"
	"golang-final-project/Models/Response"
	"golang-final-project/Models/Users"
	"golang-final-project/Test"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetUsers_OK(t *testing.T) {
	db := Test.SetUpTestDB()

	// db.Exec("DELETE FROM users WHERE id=(SELECT MAX(id) FROM users)")
	db.Exec("ALTER TABLE users AUTO_INCREMENT = 1;")

	req, w := setGetUsersRouter(db)

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

func Test_CreateUser_OK(t *testing.T) {
	db := Test.SetUpTestDB()

	a := assert.New(t)
	user := Users.User{
		NamaDepan:      "Mohammed",
		NamaBelakang:   "Salah",
		NIK:            "3777098954678908",
		TanggalLahir:   "1992-10-25",
		NoTelepon:      "089021425378",
		Alamat:         "Jl. Melati No.11",
		Kabupaten_Kota: "Jakarta Timur",
		Provinsi:       "Jakarta",
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
		Code:    http.StatusCreated,
		Message: "successful create data",
	}
	a.Equal(expected, actual)
}

func Test_UpdateUser_OK(t *testing.T) {
	db := Test.SetUpTestDB()

	a := assert.New(t)
	user := Users.User{
		NamaDepan:      "Mohammed",
		NamaBelakang:   "Salah",
		NIK:            "3777098954678908",
		TanggalLahir:   "1992-10-25",
		NoTelepon:      "089021425378",
		Alamat:         "Jl. Melati No.11 Kp. Babakan",
		Kabupaten_Kota: "Jakarta Timur",
		Provinsi:       "Jakarta",
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
	db := Test.SetUpTestDB()

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

	db.Exec("ALTER TABLE users AUTO_INCREMENT = 1;")
}
