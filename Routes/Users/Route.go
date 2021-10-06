package UsersRoute

import (
	"golang-final-project/Configs"
	"golang-final-project/Controllers/UserHandler"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UsersRouter(route *echo.Group) {
	jwtSecretKey := os.Getenv("SECRET_JWT")
	jwt := middleware.JWT([]byte(jwtSecretKey))

	api := &UserHandler.APIEnv{DB: Configs.DB}

	route.POST("users", api.CreateUser)
	route.GET("users", api.GetUsers, jwt)
	route.GET("users/:id", api.GetUser, jwt)
	route.PUT("users/:id", api.UpdateUser, jwt)
	route.DELETE("users/:id", api.DeleteUser, jwt)
}
