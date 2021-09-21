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
		return c.JSON(http.StatusInternalServerError, Response.Respond(http.StatusInternalServerError, "Error save data to database", nil))
	}

	return c.JSON(http.StatusCreated, Response.Respond(http.StatusCreated, "Successful create data", &users))
}

//GetUsers ... Get all users
func GetAllUser(c echo.Context) error {
	var users = []Users.User{}

	result := Configs.DB.Find(&users)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.Respond(http.StatusInternalServerError, "Cannot retrieve data from database", nil))
	}

	return c.JSON(http.StatusOK, Response.Respond(http.StatusOK, "Successful retrieve data", &users))
}

//GetUserByID ... Get the user by id
func GetUserByID(c echo.Context) error {
	var user Users.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.Respond(http.StatusUnprocessableEntity, "Path parameter invalid", nil))
	}

	result := Configs.DB.First(&user, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.Respond(http.StatusInternalServerError, "Cannot retrieve data from database", nil))
	}

	return c.JSON(http.StatusOK, Response.Respond(http.StatusOK, "Successful retrieve data", &user))
}

//UpdateUser ... Update the user information
func UpdateUser(c echo.Context) error {
	var user Users.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.Respond(http.StatusUnprocessableEntity, "Path parameter invalid", nil))
	}

	result := Configs.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, Response.Respond(http.StatusNotAcceptable, "Data not Found", nil))
	}

	c.Bind(&user)
	result = Configs.DB.Save(&user)
	if result.Error != nil {
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, Response.Respond(http.StatusInternalServerError, "Cannot Update data to database", nil))
		}
	}

	return c.JSON(http.StatusAccepted, Response.Respond(http.StatusAccepted, "Successful update data", &user))
}

//DeleteUser ... Delete the user
func DeleteUser(c echo.Context) error {
	var user Users.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.Respond(http.StatusUnprocessableEntity, "Path parameter invalid", nil))
	}

	result := Configs.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, Response.Respond(http.StatusNotAcceptable, "Data not Found", nil))
	}

	result = Configs.DB.Where("id = ?", id).Unscoped().Delete(&user)
	if result.Error != nil {
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, Response.Respond(http.StatusInternalServerError, "Cannot Delete data to database", nil))
		}
	}

	return c.JSON(http.StatusAccepted, Response.Respond(http.StatusAccepted, "Successful Delete data", &user))
}
