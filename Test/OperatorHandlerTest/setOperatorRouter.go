package OperatorHandlerTest

import (
	"bytes"
	"golang-final-project/Controllers/OperatorSampahHandler"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func setCreateOperatorRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &OperatorSampahHandler.APIEnv{DB: db}
	e.POST("/api/v1/operator-sampah", api.CreateOperatorSampah)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/operator-sampah", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}

func setGetOperatorsRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	api := &OperatorSampahHandler.APIEnv{DB: db}
	e.GET("/api/v1/operator-sampah", api.GetAllOperatorSampah)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/operator-sampah", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w
}

func setUpdateOperatorRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &OperatorSampahHandler.APIEnv{DB: db}
	e.PUT("/api/v1/operator-sampah/:id", api.UpdateOperatorSampah)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/operator-sampah/2", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}

func setDeleteOperatorRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &OperatorSampahHandler.APIEnv{DB: db}
	e.DELETE("/api/v1/operator-sampah/:id", api.DeleteOperatorSampah)
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/operator-sampah/2", nil)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}
