package BankSampahTest

import (
	"bytes"
	"golang-final-project/Controllers/BankSampahHandler"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func setCreateBankRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &BankSampahHandler.APIEnv{DB: db}
	e.POST("/api/v1/bank-sampah", api.BankSampahRegister)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/bank-sampah", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}

func setGetBankSampahRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	api := &BankSampahHandler.APIEnv{DB: db}
	e.GET("/api/v1/bank-sampah", api.GetAllBankSampah)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/bank-sampah", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w
}

func setUpdateBankSampahRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &BankSampahHandler.APIEnv{DB: db}
	e.PUT("/api/v1/bank-sampah/:id", api.UpdateBankSampah)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/bank-sampah/2", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}

func setDeleteBankSampahRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &BankSampahHandler.APIEnv{DB: db}
	e.DELETE("/api/v1/bank-sampah/:id", api.DeleteBankSampah)
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/bank-sampah/2", nil)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}
