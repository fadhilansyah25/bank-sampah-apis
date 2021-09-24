package TransactionRoute

import (
	"golang-final-project/Controllers"

	"github.com/labstack/echo/v4"
)

func TransactionRouter(route *echo.Group) {
	// jwtSecretKey := os.Getenv("SECRET_JWT")

	// jwt := middleware.JWT([]byte(jwtSecretKey))

	route.POST("transaction", Controllers.AddTransaction)
	route.GET("transaction", Controllers.GetAllTransaction)
	route.GET("transaction/:id", Controllers.GetTransactionById)
	route.PUT("transaction/:id", Controllers.UpdateTansaction)
}
