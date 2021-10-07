package BankSampahRoute

import (
	"golang-final-project/Configs/Database"
	"golang-final-project/Controllers/BankSampahHandler"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BankSampahRouter(route *echo.Group) {
	jwtSecretKey := os.Getenv("SECRET_JWT")
	jwt := middleware.JWT([]byte(jwtSecretKey))

	api := BankSampahHandler.APIEnv{DB: Database.DB}

	route.POST("bank-sampah", api.BankSampahRegister, jwt)
	route.GET("bank-sampah", api.GetAllBankSampah, jwt)
	route.GET("bank-sampah/:id", api.GetBankSampahById, jwt)
	route.PUT("bank-sampah/:id", api.UpdateBankSampah, jwt)
	route.DELETE("bank-sampah/:id", api.DeleteBankSampah, jwt)
}
