package UsersRoute

import (
	"golang-final-project/Controllers"

	"github.com/labstack/echo/v4"
)

func UsersRouter(route *echo.Group) {
	route.POST("users", Controllers.UserRegister)
	route.GET("users", Controllers.GetAllUser)
	route.GET("users/:id", Controllers.GetUserByID)
	route.PUT("users/:id", Controllers.UpdateUser)
	route.DELETE("users/:id", Controllers.DeleteUser)
}
