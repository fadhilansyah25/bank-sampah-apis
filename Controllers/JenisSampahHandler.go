package Controllers

import (
	"golang-final-project/Configs"
	"golang-final-project/Models/Response"
	"golang-final-project/Models/Transaction"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Add Jenis Sampah
func AddJenisSampah(c echo.Context) error {
	var jenisSampah Transaction.JenisSampah
	c.Bind(&jenisSampah)

	result := Configs.DB.Create(&jenisSampah)
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
		Data:    &jenisSampah,
	})
}

// Get All Data Jenis Sampah
func GetAllJenisSampah(c echo.Context) error {
	var jenisSampah = []Transaction.JenisSampah{}

	result := Configs.DB.Find(&jenisSampah)
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
		Data:    &jenisSampah,
	})
}

// Get Jenis Sampah by ID
func GetJenisSampahById(c echo.Context) error {
	var jenisSampah Transaction.JenisSampah

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := Configs.DB.First(&jenisSampah, id)

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
		Data:    &jenisSampah,
	})
}

// Update Jenis Sampah
func UpdateJenisSampah(c echo.Context) error {
	var JenisSampah Transaction.JenisSampah

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := Configs.DB.First(&JenisSampah, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, Response.BaseResponse{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	c.Bind(&JenisSampah)
	result = Configs.DB.Save(&JenisSampah)
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
		Data:    &JenisSampah,
	})
}

// Delete Jenis Sampah
func DeleteJenisSampah(c echo.Context) error {
	var jenisSampah Transaction.JenisSampah

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := Configs.DB.First(&jenisSampah, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, Response.BaseResponse{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	result = Configs.DB.Where("id = ?", id).Unscoped().Delete(&jenisSampah)
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
		Data:    &jenisSampah,
	})
}
