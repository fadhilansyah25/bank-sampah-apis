package UserHandler

import (
	"fmt"
	"golang-final-project/Driver/UserDriver"
	"golang-final-project/Models/Response"
	"golang-final-project/Models/Users"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) GetUsers(c echo.Context) error {
	users, err := UserDriver.GetAllUsers(a.DB)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, Response.BaseResponse{
		Code:    http.StatusOK,
		Message: "successful retrieve data",
		Data:    &users,
	})
}

func (a *APIEnv) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, exists, err := UserDriver.GetUserByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cannot retrieve data from database",
			Data:    nil,
		})
	}

	if !exists {
		return c.JSON(http.StatusNotFound, Response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "data not found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, Response.BaseResponse{
		Code:    http.StatusOK,
		Message: "successful retrieve data",
		Data:    &user,
	})
}

func (a *APIEnv) CreateUser(c echo.Context) error {
	user := Users.User{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "error process header",
			Data:    nil,
		})
	}

	if err := UserDriver.CreateUser(a.DB, &user); err != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "error save data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, Response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "successful create data",
		Data:    &user,
	})
}

func (a *APIEnv) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	_, exists, err := UserDriver.GetUserByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "error process to database",
			Data:    nil,
		})
	}

	if !exists {
		return c.JSON(http.StatusNotFound, Response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "record not exist",
			Data:    nil,
		})
	}

	err = UserDriver.DeleteUser(id, a.DB)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cannot delete data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "successful delete data",
		Data:    nil,
	})
}

func (a *APIEnv) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user, exists, err := UserDriver.GetUserByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "path parameter invalid",
			Data:    nil,
		})
	}

	if !exists {
		return c.JSON(http.StatusNotFound, Response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "record not exist",
			Data:    nil,
		})
	}

	err = c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "cannot process data",
			Data:    nil,
		})
	}

	if err := UserDriver.UpdateUser(a.DB, &user, id); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cannot update data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "successful update data",
		Data:    user,
	})
}
