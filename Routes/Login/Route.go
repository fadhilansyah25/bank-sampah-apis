package LoginRoute

import (
	"golang-final-project/Controllers"

	"github.com/labstack/echo/v4"
)

func UserLoginRoute(route *echo.Group) {
	route.POST("user-login", Controllers.CreateUserLogin)
	route.GET("user-login", Controllers.GetAllUserLogin)
	route.POST("login", Controllers.UserLogin)
	route.PUT("user-login/:id", Controllers.UpdateUserLogin)
}
