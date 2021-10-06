package JenisSampahRoute

import (
	"golang-final-project/Configs"
	"golang-final-project/Controllers/JenisSampahHandler"

	"github.com/labstack/echo/v4"
)

func JenisSampahRouter(route *echo.Group) {
	// jwtSecretKey := os.Getenv("SECRET_JWT")
	// jwt := middleware.JWT([]byte(jwtSecretKey))

	api := &JenisSampahHandler.APIEnv{DB: Configs.DB}

	route.POST("jenis-sampah", api.AddJenisSampah)
	route.GET("jenis-sampah", api.GetAllJenisSampah)
	route.GET("jenis-sampah/:id", api.GetJenisSampahById)
	route.PUT("jenis-sampah/:id", api.UpdateJenisSampah)
	route.DELETE("jenis-sampah/:id", api.DeleteJenisSampah)
}
