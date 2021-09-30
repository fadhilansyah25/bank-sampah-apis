package Controllers

import (
	"golang-final-project/Configs"
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

	res := Configs.DB.Create(&Transaction.Transaction{
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
			Message: "Cannot save data to database",
			Data:    nil,
		})
	}

	var trans = Transaction.Transaction{}

	Configs.DB.Preload("DetailTransaction").Last(&trans)

	return c.JSON(http.StatusCreated, Response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Successful create data",
		Data:    &trans,
	})
}

// Get All Data Transaction
func GetAllTransaction(c echo.Context) error {
	var trans = []Transaction.Transaction{}

	result := Configs.DB.Preload("DetailTransaction").Find(&trans)
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
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := Configs.DB.Preload("DetailTransaction").First(&trans, id)

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
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := Configs.DB.Preload("DetailTransaction").First(&transaction, id)
	if result.Error != nil {
		return c.JSON(http.StatusGone, Response.BaseResponse{
			Code:    http.StatusGone,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	c.Bind(&transaction)
	result = Configs.DB.Save(&transaction)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Cannot Update data to database",
			Data:    nil,
		})
	}

	Configs.DB.Preload("DetailTransaction").First(&transaction, id)
	return c.JSON(http.StatusAccepted, Response.BaseResponse{
		Code:    http.StatusAccepted,
		Message: "Successful update data",
		Data:    &transaction,
	})
}
