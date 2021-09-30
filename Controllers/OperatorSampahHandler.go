package Controllers

import (
	"golang-final-project/Configs"
	"golang-final-project/Models/BankSampah"
	"golang-final-project/Models/Response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Register Operator Sampah
func OperatorSampahRegister(c echo.Context) error {
	var operator BankSampah.OperatorSampah
	c.Bind(&operator)

	result := Configs.DB.Create(&operator)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error save data to database",
			Data:    nil,
		})
	}

	if res := Configs.DB.Preload("BankSampah").Find(&operator); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error retrieve data after saved",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, Response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Successful create data",
		Data:    &operator,
	})
}

// Get All Operator Sampah Data
func GetAllOperatorSampah(c echo.Context) error {
	var operator = []BankSampah.OperatorSampah{}

	result := Configs.DB.Preload("BankSampah").Find(&operator)
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
		Data:    &operator,
	})
}

// Get Bank Sampah by ID
func GetOperatorSampahById(c echo.Context) error {
	var operator BankSampah.OperatorSampah

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := Configs.DB.Preload("BankSampah").First(&operator, id)
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
		Data:    &operator,
	})
}

// Update Bank Sampah
func UpdateOperatorSampah(c echo.Context) error {
	var operator BankSampah.OperatorSampah

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := Configs.DB.First(&operator, id)
	if result.Error != nil {
		return c.JSON(http.StatusGone, Response.BaseResponse{
			Code:    http.StatusGone,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	c.Bind(&operator)
	Configs.DB.Model(&operator).Update("BankSampahId", operator.BankSampahId)
	result = Configs.DB.Save(&operator)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Cannot Update data to database",
			Data:    nil,
		})
	}

	Configs.DB.Preload("BankSampah").First(&operator, id)
	return c.JSON(http.StatusAccepted, Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "Successful update data",
		Data:    &operator,
	})
}

// Delete Operator Sampah
func DeleteOperatorSampah(c echo.Context) error {
	var operator BankSampah.OperatorSampah

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := Configs.DB.First(&operator, id)
	if result.Error != nil {
		return c.JSON(http.StatusGone, Response.BaseResponse{
			Code:    http.StatusGone,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	result = Configs.DB.Where("id = ?", id).Unscoped().Delete(&operator)
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
		Data:    &operator,
	})
}
