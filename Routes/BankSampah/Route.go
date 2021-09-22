package BankSampahRoute

import (
	"golang-final-project/Controllers"

	"github.com/labstack/echo/v4"
)

func BankSampahRouter(route *echo.Group) {
	// jwtSecretKey := os.Getenv("SECRET_JWT")

	// jwt := middleware.JWT([]byte(jwtSecretKey))

	route.POST("bank-sampah", Controllers.BankSampahRegister)
	route.GET("bank-sampah", Controllers.GetAllBankSampah)
	route.GET("bank-sampah/:id", Controllers.GetBankSampahById)
	route.PUT("bank-sampah/:id", Controllers.UpdateBankSampah)
	route.DELETE("bank-sampah/:id", Controllers.DeleteBankSampah)
}
