package UsersRoute

import (
	"golang-final-project/Controllers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UsersRouter(route *echo.Group) {
	jwtSecretKey := os.Getenv("SECRET_JWT")

	jwt := middleware.JWT([]byte(jwtSecretKey))

	route.POST("users", Controllers.UserRegister, jwt)
	route.GET("users", Controllers.GetAllUser)
	route.GET("users/:id", Controllers.GetUserByID, jwt)
	route.PUT("users/:id", Controllers.UpdateUser, jwt)
	route.DELETE("users/:id", Controllers.DeleteUser, jwt)
}
