package Controllers

import (
	"golang-final-project/Configs"
	"golang-final-project/Models/Response"
	"golang-final-project/Models/Users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UserRegister ... Post the new user
func UserRegister(c echo.Context) error {
	var users Users.User
	c.Bind(&users)

	result := Configs.DB.Create(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error save data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, Response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Successful create data",
		Data:    &users,
	})
}

//GetUsers ... Get all users
func GetAllUser(c echo.Context) error {
	var users = []Users.User{}

	result := Configs.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusGone, Response.BaseResponse{
			Code:    http.StatusGone,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, Response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successful retrieve data",
		Data:    &users,
	})
}

//GetUserByID ... Get the user by id
func GetUserByID(c echo.Context) error {
	var user Users.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := Configs.DB.First(&user, id)

	if result.Error != nil {
		return c.JSON(http.StatusGone, Response.BaseResponse{
			Code:    http.StatusGone,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, Response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successful retrieve data",
		Data:    &user,
	})
}

//UpdateUser ... Update the user information
func UpdateUser(c echo.Context) error {
	var user Users.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := Configs.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, Response.BaseResponse{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	c.Bind(&user)
	result = Configs.DB.Save(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Cannot Update data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "Successful update data",
		Data:    &user,
	})
}

//DeleteUser ... Delete the user
func DeleteUser(c echo.Context) error {
	var user Users.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := Configs.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, Response.BaseResponse{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	result = Configs.DB.Where("id = ?", id).Unscoped().Delete(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Cannot Delete data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "Successful Delete data",
		Data:    &user,
	})
}
