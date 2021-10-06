package UserHandlerTest

import (
	"bytes"
	"golang-final-project/Controllers/UserHandler"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func setGetUsersRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	api := &UserHandler.APIEnv{DB: db}
	e.GET("/api/v1/users", api.GetUsers)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/users", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w
}

func setCreateUserRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &UserHandler.APIEnv{DB: db}
	e.POST("/api/v1/users", api.CreateUser)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/users", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}

func setUpdateUserRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &UserHandler.APIEnv{DB: db}
	e.PUT("/api/v1/users/:id", api.UpdateUser)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/users/2", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}

func setDeleteUserRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder, error) {
	e := echo.New()
	api := &UserHandler.APIEnv{DB: db}
	e.DELETE("/api/v1/users/:id", api.DeleteUser)
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/users/2", nil)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	return req, w, nil
}
