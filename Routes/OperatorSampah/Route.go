package OperatorSampahRoute

import (
	"golang-final-project/Configs/Database"
	"golang-final-project/Controllers/OperatorSampahHandler"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func OperatorSampahRoute(route *echo.Group) {
	jwtSecretKey := os.Getenv("SECRET_JWT")
	jwt := middleware.JWT([]byte(jwtSecretKey))

	api := &OperatorSampahHandler.APIEnv{DB: Database.DB}

	route.POST("operator-sampah", api.CreateOperatorSampah)
	route.GET("operator-sampah", api.GetAllOperatorSampah, jwt)
	route.GET("operator-sampah/:id", api.GetOperatorSampahByID, jwt)
	route.PUT("operator-sampah/:id", api.UpdateOperatorSampah, jwt)
	route.DELETE("operator-sampah/:id", api.DeleteOperatorSampah, jwt)
}
