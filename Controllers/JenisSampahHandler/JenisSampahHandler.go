package JenisSampahHandler

import (
	"fmt"
	"golang-final-project/Driver/JenisSampahDriver"
	"golang-final-project/Models/Response"
	"golang-final-project/Models/Transaction"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type APIEnv struct {
	DB *gorm.DB
}

// Add Jenis Sampah
func (a *APIEnv) AddJenisSampah(c echo.Context) error {
	var jenisSampah Transaction.JenisSampah
	err := c.Bind(&jenisSampah)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "error process header",
			Data:    nil,
		})
	}

	if err := JenisSampahDriver.CreateJenisSampah(a.DB, &jenisSampah); err != nil {
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
		Data:    &jenisSampah,
	})
}

// Get All Data Jenis Sampah
func (a *APIEnv) GetAllJenisSampah(c echo.Context) error {
	jenisSampah, err := JenisSampahDriver.GetAllJenisSampah(a.DB)
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
		Data:    &jenisSampah,
	})
}

// Get Jenis Sampah by ID
func (a *APIEnv) GetJenisSampahById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "path parameter invalid",
			Data:    nil,
		})
	}

	jenisSampah, exists, err := JenisSampahDriver.GetJenisSampahByID(fmt.Sprint(id), a.DB)
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
		Message: "Successful retrieve data",
		Data:    &jenisSampah,
	})
}

// Update Jenis Sampah
func (a *APIEnv) UpdateJenisSampah(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "path parameter invalid",
			Data:    nil,
		})
	}

	jenisSampah, exists, err := JenisSampahDriver.GetJenisSampahByID(fmt.Sprint(id), a.DB)
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

	err = c.Bind(&jenisSampah)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "cannot process data",
			Data:    nil,
		})
	}

	if err := JenisSampahDriver.UpdateJenisSampah(a.DB, &jenisSampah, fmt.Sprint(id)); err != nil {
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
		Data:    &jenisSampah,
	})
}

// Delete Jenis Sampah
func (a *APIEnv) DeleteJenisSampah(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	_, exists, err := JenisSampahDriver.GetJenisSampahByID(fmt.Sprint(id), a.DB)
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

	err = JenisSampahDriver.DeleteJenisSampah(fmt.Sprint(id), a.DB)
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
