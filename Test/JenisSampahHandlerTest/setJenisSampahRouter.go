package JenisSampahTest

import (
	"bytes"
	"golang-final-project/Controllers/JenisSampahHandler"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func setCreateJenisSampahRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &JenisSampahHandler.APIEnv{DB: db}
	e.POST("/api/v1/jenis-sampah", api.AddJenisSampah)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/jenis-sampah", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}

func setGetAllJenisSampahRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	api := &JenisSampahHandler.APIEnv{DB: db}
	e.GET("/api/v1/jenis-sampah", api.GetAllJenisSampah)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/jenis-sampah", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w
}

func setUpdateJenisSampahRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &JenisSampahHandler.APIEnv{DB: db}
	e.PUT("/api/v1/jenis-sampah/:id", api.UpdateJenisSampah)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/jenis-sampah/1", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}

func setDeleteJenisSampahRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &JenisSampahHandler.APIEnv{DB: db}
	e.DELETE("/api/v1/jenis-sampah/:id", api.DeleteJenisSampah)
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/jenis-sampah/1", nil)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}
