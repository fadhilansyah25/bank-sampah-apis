package BankSampahHandler

import (
	"fmt"
	"golang-final-project/Driver/BankSampahDriver"
	"golang-final-project/Models/BankSampah"
	"golang-final-project/Models/Response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type APIEnv struct {
	DB *gorm.DB
}

// Register Bank Sampah
func (a *APIEnv) BankSampahRegister(c echo.Context) error {
	var bankSampah BankSampah.BankSampah
	c.Bind(&bankSampah)

	err := BankSampahDriver.CreateBankSampah(a.DB, &bankSampah)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "error save data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, Response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "successful create data",
		Data:    &bankSampah,
	})
}

// Get All Bank Sampah Data
func (a *APIEnv) GetAllBankSampah(c echo.Context) error {
	bankSampah, err := BankSampahDriver.GetAllBankSampah(a.DB)
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
		Data:    &bankSampah,
	})
}

// Get Bank Sampah by ID
func (a *APIEnv) GetBankSampahById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "path parameter invalid",
			Data:    nil,
		})
	}

	bankSampah, exists, err := BankSampahDriver.GetBankSampahByID(fmt.Sprint(id), a.DB)
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
		Data:    &bankSampah,
	})
}

// Update Bank Sampah
func (a *APIEnv) UpdateBankSampah(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "path parameter invalid",
			Data:    nil,
		})
	}

	bankSampah, exists, err := BankSampahDriver.GetBankSampahByID(fmt.Sprint(id), a.DB)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotAcceptable, Response.BaseResponse{
			Code:    http.StatusNotAcceptable,
			Message: "data not found",
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

	c.Bind(&bankSampah)
	if err := BankSampahDriver.UpdateBankSampah(a.DB, &bankSampah, fmt.Sprint(id)); err != nil {
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
		Data:    &bankSampah,
	})
}

func (a *APIEnv) DeleteBankSampah(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "path parameter invalid",
			Data:    nil,
		})
	}

	_, exists, err := BankSampahDriver.GetBankSampahByID(fmt.Sprint(id), a.DB)
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

	err = BankSampahDriver.DeleteBankSampah(fmt.Sprint(id), a.DB)
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
