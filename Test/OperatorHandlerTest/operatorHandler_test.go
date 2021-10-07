package OperatorHandlerTest

import (
	"bytes"
	"encoding/json"
	"golang-final-project/Models/BankSampah"
	"golang-final-project/Models/Response"
	"golang-final-project/Test"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateOperator_OK(t *testing.T) {
	db := Test.SetUpTestDB()

	// db.Exec("DELETE FROM operator_sampahs WHERE id=(SELECT MAX(id) FROM operator_sampahs)")
	db.Exec("ALTER TABLE operator_sampahs AUTO_INCREMENT = 1;")

	a := assert.New(t)
	operator := BankSampah.OperatorSampah{
		NIK:            "801929938847799293",
		BankSampahId:   1,
		NamaDepan:      "Bambang",
		NamaBelakang:   "Rianto",
		TanggalLahir:   "1983-02-03",
		NoTelepon:      "0897761627839",
		Alamat:         "Jl. Lestari Bersama 4, No. 102 Kp. Angke",
		Kabupaten_Kota: "Jakarta Utara",
		Provinsi:       "DKI Jakarta",
	}

	reqBody, err := json.Marshal(operator)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateOperatorRouter(db, bytes.NewBuffer(reqBody))
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

func Test_GetOperators_OK(t *testing.T) {
	db := Test.SetUpTestDB()

	req, w := setGetOperatorsRouter(db)

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

func Test_UpdateOperator_OK(t *testing.T) {
	db := Test.SetUpTestDB()

	a := assert.New(t)
	operator := BankSampah.OperatorSampah{
		NIK:            "801929938847799293",
		BankSampahId:   1,
		NamaDepan:      "Bambang",
		NamaBelakang:   "Pamungkas",
		TanggalLahir:   "1983-02-03",
		NoTelepon:      "0897761627839",
		Alamat:         "Jl. Lestari segar 4, No. 102 Kp. Angke",
		Kabupaten_Kota: "Jakarta Utara",
		Provinsi:       "DKI Jakarta",
	}

	reqBody, err := json.Marshal(operator)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setUpdateOperatorRouter(db, bytes.NewBuffer(reqBody))
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

func Test_DeleteOperator_OK(t *testing.T) {
	db := Test.SetUpTestDB()

	a := assert.New(t)

	req, w, err := setDeleteOperatorRouter(db)
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
