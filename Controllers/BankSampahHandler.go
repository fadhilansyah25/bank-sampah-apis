package Controllers

import (
	"golang-final-project/Configs"
	"golang-final-project/Models/BankSampah"
	"golang-final-project/Models/Response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Register Bank Sampah
func BankSampahRegister(c echo.Context) error {
	var bankSampah BankSampah.BankSampah
	c.Bind(&bankSampah)

	result := Configs.DB.Create(&bankSampah)
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
		Data:    &bankSampah,
	})
}

// Get All Bank Sampah Data
func GetAllBankSampah(c echo.Context) error {
	var bankSampah = []BankSampah.BankSampah{}

	result := Configs.DB.Find(&bankSampah)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, Response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successful retrieve data",
		Data:    &bankSampah,
	})
}

// Get Bank Sampah by ID
func GetBankSampahById(c echo.Context) error {
	var bankSampah BankSampah.BankSampah

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := Configs.DB.First(&bankSampah, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, Response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successful retrieve data",
		Data:    &bankSampah,
	})
}

// Update Bank Sampah
func UpdateBankSampah(c echo.Context) error {
	var bankSampah BankSampah.BankSampah

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := Configs.DB.First(&bankSampah, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, Response.BaseResponse{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	c.Bind(&bankSampah)
	result = Configs.DB.Save(&bankSampah)
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
		Data:    &bankSampah,
	})
}

// Delete Bank Sampah
func DeleteBankSampah(c echo.Context) error {
	var bankSampah BankSampah.BankSampah

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := Configs.DB.First(&bankSampah, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, Response.BaseResponse{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	result = Configs.DB.Where("id = ?", id).Unscoped().Delete(&bankSampah)
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
		Data:    &bankSampah,
	})
}
