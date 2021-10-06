package UserLoginHandler

import (
	"fmt"
	"golang-final-project/Driver/UserDriver"
	"golang-final-project/Driver/UserLoginDriver"
	"golang-final-project/Helper"
	"golang-final-project/Middleware"
	"golang-final-project/Models/Response"
	"golang-final-project/Models/UserLogins"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type APIEnv struct {
	DB *gorm.DB
}

// Create user login ...
func (a *APIEnv) CreateUserLogin(c echo.Context) error {
	var userlogin UserLogins.LoginDataUsers

	// binding data from header request
	err := c.Bind(&userlogin)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "error process header",
			Data:    nil,
		})
	}

	// get data user from database
	user, exists, err := UserDriver.GetUserByID(fmt.Sprint(userlogin.UserId), a.DB)
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

	userlogin.Password = Helper.Encript(userlogin.Password)
	userlogin.UserId = user.Id

	err = UserLoginDriver.CreateUserLogin(a.DB, &userlogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "failed create data to database",
			Data:    nil,
		})
	}

	userlogin, _, _ = UserLoginDriver.GetUserLoginByID(fmt.Sprint(userlogin.UserId), a.DB)
	return c.JSON(http.StatusCreated, Response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "successful create data",
		Data:    &userlogin,
	})
}

// Login handler ...
func (a *APIEnv) Login(c echo.Context) error {
	var login UserLogins.Login
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

	var userlogin UserLogins.LoginDataUsers
	userlogin, err := UserLoginDriver.Login(a.DB, &login, &userlogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cannot retrieve data from database",
			Data:    nil,
		})
	}

	// check if password, username, or email request is not match in database
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

// get all user with validatation data ...
func (a *APIEnv) GetAlluserLogin(c echo.Context) error {
	userlogins, err := UserLoginDriver.GetAllUserLogins(a.DB)
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
		Data:    &userlogins,
	})
}

// get user login by id handler ...
func (a *APIEnv) GetUserLoginByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "path parameter invalid",
			Data:    nil,
		})
	}

	user, exists, err := UserLoginDriver.GetUserLoginByID(fmt.Sprint(id), a.DB)
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

func (a *APIEnv) UpdateUserLogin(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "path parameter invalid",
			Data:    nil,
		})
	}

	userlogin, exists, err := UserLoginDriver.GetUserLoginByID(fmt.Sprint(id), a.DB)
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

	err = c.Bind(&userlogin)
	userlogin.Password = Helper.Encript(userlogin.Password)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "cannot process data",
			Data:    nil,
		})
	}

	if err := UserLoginDriver.UpdateUserLogin(a.DB, &userlogin); err != nil {
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
		Data:    &userlogin,
	})
}
