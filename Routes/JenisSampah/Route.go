package JenisSampahRoute

import (
	"golang-final-project/Configs/Database"
	"golang-final-project/Controllers/JenisSampahHandler"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JenisSampahRouter(route *echo.Group) {
	jwtSecretKey := os.Getenv("SECRET_JWT")
	jwt := middleware.JWT([]byte(jwtSecretKey))

	api := &JenisSampahHandler.APIEnv{DB: Database.DB}

	route.POST("jenis-sampah", api.AddJenisSampah, jwt)
	route.GET("jenis-sampah", api.GetAllJenisSampah, jwt)
	route.GET("jenis-sampah/:id", api.GetJenisSampahById, jwt)
	route.PUT("jenis-sampah/:id", api.UpdateJenisSampah, jwt)
	route.DELETE("jenis-sampah/:id", api.DeleteJenisSampah, jwt)
}
