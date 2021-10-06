package TransactionRoute

import (
	"golang-final-project/Controllers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TransactionRouter(route *echo.Group) {
	jwtSecretKey := os.Getenv("SECRET_JWT")
	jwt := middleware.JWT([]byte(jwtSecretKey))

	route.POST("transaction", Controllers.AddTransaction, jwt)
	route.GET("transaction", Controllers.GetAllTransaction, jwt)
	route.GET("transaction/:id", Controllers.GetTransactionById, jwt)
	route.PUT("transaction/:id", Controllers.UpdateTansaction, jwt)
}
