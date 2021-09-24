package LoginRoute

import (
	"golang-final-project/Controllers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserLoginRoute(route *echo.Group) {
	jwtSecretKey := os.Getenv("SECRET_JWT")
	jwt := middleware.JWT([]byte(jwtSecretKey))

	route.POST("create-login", Controllers.CreateUserLogin, jwt)
	route.GET("user-login", Controllers.GetAllUserLogin, jwt)
	route.POST("login", Controllers.UserLogin)
	route.PUT("user-login/:id", Controllers.UpdateUserLogin, jwt)
}
