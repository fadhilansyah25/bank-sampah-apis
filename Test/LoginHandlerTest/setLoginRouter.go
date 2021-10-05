package LoginHandlerTest

import (
	"bytes"
	"golang-final-project/Controllers/UserLoginHandler"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func setCreateUserLoginRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &UserLoginHandler.APIEnv{DB: db}
	e.POST("/api/v1/create-login", api.CreateUserLogin)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/create-login", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}

func setLoginRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &UserLoginHandler.APIEnv{DB: db}
	e.POST("/api/v1/login", api.Login)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/login", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}

func setUpdateLoginRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &UserLoginHandler.APIEnv{DB: db}
	e.POST("/api/v1/user-login/:id", api.UpdateUserLogin)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/user-login/1", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}

func setGetUserLoginByIDRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	api := &UserLoginHandler.APIEnv{DB: db}
	e.GET("/api/user-login/:id", api.GetUserLoginByID)
	req, err := http.NewRequest(http.MethodGet, "/api/user-login/1", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w
}
