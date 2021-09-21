package Controllers

import (
	"golang-final-project/Configs"
	"golang-final-project/Helper"
	"golang-final-project/Models/Login"
	"golang-final-project/Models/Response"
	"golang-final-project/Models/Users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateUserLogin ... Post new meta Login
func CreateUserLogin(c echo.Context) error {
	var userlogin Login.LoginDataUsers
	var user Users.User

	c.Bind(&userlogin)
	if res := Configs.DB.Where("id = ?", userlogin.UserId).Find(&user); res.Error != nil {
		return c.JSON(http.StatusNotAcceptable, Response.Respond(http.StatusNotAcceptable, "Data not Found", nil))
	}

	userlogin.Password = Helper.Encript(userlogin.Password)
	userlogin.User = user

	if res := Configs.DB.Create(&userlogin); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.Respond(http.StatusInternalServerError, "Failed create data to database", nil))
	}

	return c.JSON(http.StatusCreated, Response.Respond(http.StatusCreated, "Success", &userlogin))
}

// GetAllUserVerification ... All User Verifation Data
func GetAllUserLogin(c echo.Context) error {
	var userlogins = []Login.LoginDataUsers{}

	if res := Configs.DB.Joins("User").Find(&userlogins); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.Respond(http.StatusInternalServerError, "Cannot retrieve data from database", nil))
	}

	return c.JSON(http.StatusOK, Response.Respond(http.StatusOK, "Successful retrieve data", &userlogins))
}

// UserLoginValidation ... All User Verifation Data
func UserLogin(c echo.Context) error {
	var login Login.Login
	c.Bind(&login)

	if login.Email == "" && login.Username == "" {
		return c.JSON(http.StatusBadRequest, Response.Respond(http.StatusBadRequest, "invalid input", nil))
	}

	if login.Password == "" {
		return c.JSON(http.StatusBadRequest, Response.Respond(http.StatusBadRequest, "invalid input", nil))
	}

	var userlogin Login.LoginDataUsers
	res := Configs.DB.Where("username = ? OR email = ?", login.Username, login.Email).Find(&userlogin)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.Respond(http.StatusInternalServerError, "cannot retrieve data", &userlogin))
	}

	if userlogin.Email == "" || userlogin.Username == "" || userlogin.Password == "" {
		return c.JSON(http.StatusForbidden, Response.Respond(http.StatusForbidden, "access denied: email or username invalid", nil))
	}

	access := Helper.ComparePassword(userlogin.Password, login.Password)
	if !access {
		return c.JSON(http.StatusForbidden, Response.Respond(http.StatusForbidden, "access denied: password did not match", &access))
	}

	return c.JSON(http.StatusAccepted, Response.Respond(http.StatusAccepted, "access complete", &access))
}

// UpdateUserLogin ... Update User Login Data
func UpdateUserLogin(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.Respond(http.StatusUnprocessableEntity, "Path parameter invalid", nil))
	}

	var userlogin Login.LoginDataUsers
	result := Configs.DB.First(&userlogin, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, Response.Respond(http.StatusNotAcceptable, "Data not Found", nil))
	}

	c.Bind(&userlogin)
	userlogin.Password = Helper.Encript(userlogin.Password)

	result = Configs.DB.Save(&userlogin)
	if result.Error != nil {
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, Response.Respond(http.StatusInternalServerError, "Cannot Update data to database", nil))
		}
	}

	return c.JSON(http.StatusAccepted, Response.Respond(http.StatusAccepted, "Successful update data", &userlogin))
}
