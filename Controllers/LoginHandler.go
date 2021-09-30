package Controllers

import (
	"golang-final-project/Configs"
	"golang-final-project/Helper"
	"golang-final-project/Middleware"
	"golang-final-project/Models/Response"
	"golang-final-project/Models/Users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateUserLogin ... Post new meta Login
func CreateUserLogin(c echo.Context) error {
	var userlogin Users.LoginDataUsers
	var user Users.User

	c.Bind(&userlogin)
	if res := Configs.DB.Where("id = ?", userlogin.UserId).Find(&user); res.Error != nil {
		return c.JSON(http.StatusNotAcceptable, Response.BaseResponse{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	userlogin.Password = Helper.Encript(userlogin.Password)
	userlogin.User = user

	if res := Configs.DB.Create(&userlogin); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed create data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "Success create userlogin",
		Data:    &userlogin,
	})

}

// GetAllUserVerification ... All User Verifation Data
func GetAllUserLogin(c echo.Context) error {
	var userlogins = []Users.LoginDataUsers{}

	if res := Configs.DB.Joins("User").Find(&userlogins); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "Successful retrieve data",
		Data:    &userlogins,
	})
}

// UpdateUserLogin ... Update User Login Data
func GetUserLoginByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	var userlogin Users.LoginDataUsers
	result := Configs.DB.First(&userlogin, id)
	if result.Error != nil {
		return c.JSON(http.StatusGone, Response.BaseResponse{
			Code:    http.StatusGone,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "Successful update data",
		Data:    &userlogin,
	})
}

// UserLoginValidation ... All User Verifation Data
func UserLogin(c echo.Context) error {
	var login Users.Login
	c.Bind(&login)

	// check form input
	if login.Email == "" && login.Username == "" {
		return c.JSON(http.StatusBadRequest, Response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid input",
			Data:    nil,
		})
	} else if login.Password == "" {
		return c.JSON(http.StatusBadRequest, Response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid input",
			Data:    nil,
		})
	}

	// check data to database
	var userlogin Users.LoginDataUsers
	res := Configs.DB.Where("username = ? OR email = ?", login.Username, login.Email).Find(&userlogin)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	if userlogin.Email == "" || userlogin.Username == "" || userlogin.Password == "" {
		return c.JSON(http.StatusForbidden, Response.BaseResponse{
			Code:    http.StatusForbidden,
			Message: "access denied: email or username invalid",
			Data:    nil,
		})
	}

	// compare password with hashed
	access := Helper.ComparePassword(userlogin.Password, login.Password)
	if !access {
		return c.JSON(http.StatusForbidden, Response.BaseResponse{
			Code:    http.StatusForbidden,
			Message: "access denied: password did not match",
			Data:    nil,
		})
	}

	// get token
	tokenLogin, err := Middleware.GenerateTokenJWT(int(userlogin.UserId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "failed create token",
			Data:    nil,
		})
	}

	// set login response
	loginResponse := Response.LoginResponse{
		Id:        int(userlogin.UserId),
		Username:  userlogin.Username,
		Email:     userlogin.Email,
		Token:     tokenLogin,
		CreatedAt: userlogin.CreatedAt,
		UpdatedAt: userlogin.UpdatedAt,
	}

	return c.JSON(http.StatusAccepted, Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "access complete",
		Data:    &loginResponse,
	})
}

// UpdateUserLogin ... Update User Login Data
func UpdateUserLogin(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	var userlogin Users.LoginDataUsers
	result := Configs.DB.First(&userlogin, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, Response.BaseResponse{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	c.Bind(&userlogin)
	userlogin.Password = Helper.Encript(userlogin.Password)

	result = Configs.DB.Save(&userlogin)
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
		Data:    &userlogin,
	})
}
