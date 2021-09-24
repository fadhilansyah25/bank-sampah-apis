package OperatorSampahRoute

import (
	"golang-final-project/Controllers"

	"github.com/labstack/echo/v4"
)

func OperatorSampahRoute(route *echo.Group) {
	// jwtSecretKey := os.Getenv("SECRET_JWT")

	// jwt := middleware.JWT([]byte(jwtSecretKey))

	route.POST("operator-sampah", Controllers.OperatorSampahRegister)
	route.GET("operator-sampah", Controllers.GetAllOperatorSampah)
	route.GET("operator-sampah/:id", Controllers.GetOperatorSampahById)
	route.PUT("operator-sampah/:id", Controllers.UpdateOperatorSampah)
	route.DELETE("operator-sampah/:id", Controllers.DeleteOperatorSampah)
}
