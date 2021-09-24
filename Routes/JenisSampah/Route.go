package JenisSampahRoute

import (
	"golang-final-project/Controllers"

	"github.com/labstack/echo/v4"
)

func JenisSampahRouter(route *echo.Group) {
	// jwtSecretKey := os.Getenv("SECRET_JWT")

	// jwt := middleware.JWT([]byte(jwtSecretKey))

	route.POST("jenis-sampah", Controllers.AddJenisSampah)
	route.GET("jenis-sampah", Controllers.GetAllJenisSampah)
	route.GET("jenis-sampah/:id", Controllers.GetJenisSampahById)
	route.PUT("jenis-sampah/:id", Controllers.UpdateJenisSampah)
	route.DELETE("jenis-sampah/:id", Controllers.DeleteJenisSampah)
}
