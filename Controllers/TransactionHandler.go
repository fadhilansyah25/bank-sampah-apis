package Controllers

import (
	"golang-final-project/Configs/Database"
	"golang-final-project/Models/Response"
	"golang-final-project/Models/Transaction"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Add Transaction
func AddTransaction(c echo.Context) error {

	var transaction Transaction.TransactionReq
	c.Bind(&transaction)

	var totalQty int
	var totalPrice float32
	for _, item := range transaction.Detail {
		totalQty += item.Qty
		totalPrice += item.TotalPrice
	}

	transaction.TotalQty = totalQty
	transaction.TotalTransaction = totalPrice

	res := Database.DB.Create(&Transaction.Transaction{
		BankSampahId:      transaction.BankSampahId,
		UserId:            transaction.UserId,
		OperatorSampahId:  transaction.OperatorSampahId,
		TotalQty:          transaction.TotalQty,
		TotalTransaction:  transaction.TotalTransaction,
		Status:            transaction.Status,
		DetailTransaction: transaction.Detail,
	})

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cannot save data to database",
			Data:    nil,
		})
	}

	var trans = Transaction.Transaction{}

	Database.DB.Preload("DetailTransaction").Last(&trans)

	return c.JSON(http.StatusCreated, Response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "successful create data",
		Data:    &trans,
	})
}

// Get All Data Transaction
func GetAllTransaction(c echo.Context) error {
	var trans = []Transaction.Transaction{}

	result := Database.DB.Preload("DetailTransaction").Find(&trans)
	if result.Error != nil {
		return c.JSON(http.StatusGone, Response.BaseResponse{
			Code:    http.StatusGone,
			Message: "cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, Response.BaseResponse{
		Code:    http.StatusOK,
		Message: "successful retrieve data",
		Data:    &trans,
	})
}

// Get Jenis Sampah by ID
func GetTransactionById(c echo.Context) error {
	var trans Transaction.Transaction

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "path parameter invalid",
			Data:    nil,
		})
	}

	result := Database.DB.Preload("DetailTransaction").First(&trans, id)

	if result.Error != nil {
		return c.JSON(http.StatusGone, Response.BaseResponse{
			Code:    http.StatusGone,
			Message: "cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, Response.BaseResponse{
		Code:    http.StatusOK,
		Message: "successful retrieve data",
		Data:    &trans,
	})
}

// Update Bank Sampah
func UpdateTansaction(c echo.Context) error {
	var transaction Transaction.Transaction

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "path parameter invalid",
			Data:    nil,
		})
	}

	result := Database.DB.Preload("DetailTransaction").First(&transaction, id)
	if result.Error != nil {
		return c.JSON(http.StatusGone, Response.BaseResponse{
			Code:    http.StatusGone,
			Message: "data not found",
			Data:    nil,
		})
	}

	c.Bind(&transaction)
	result = Database.DB.Save(&transaction)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cannot update data to database",
			Data:    nil,
		})
	}

	Database.DB.Preload("DetailTransaction").First(&transaction, id)
	return c.JSON(http.StatusAccepted, Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "successful update data",
		Data:    &transaction,
	})
}
