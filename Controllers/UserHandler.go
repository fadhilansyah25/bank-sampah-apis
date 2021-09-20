package Controllers

import (
	"golang-final-project/Configs"
	"golang-final-project/Models/Response"
	"golang-final-project/Models/Users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserRegister(c echo.Context) error {
	var user Users.User
	c.Bind(&user)

	result := Configs.DB.Create(&user)
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
		Data:    user,
	})
}
