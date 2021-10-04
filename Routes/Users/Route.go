package UsersRoute

import (
	"golang-final-project/Configs"
	"golang-final-project/Controllers/UserHandler"

	"github.com/labstack/echo/v4"
)

func UsersRouter(route *echo.Group) {
	// jwtSecretKey := os.Getenv("SECRET_JWT")

	// jwt := middleware.JWT([]byte(jwtSecretKey))
	api := &UserHandler.APIEnv{}
	api.DB = Configs.DB

	route.POST("users", api.CreateUser)
	route.GET("users", api.GetUsers)
	route.GET("users/:id", api.GetUser)
	route.PUT("users/:id", api.UpdateUser)
	route.DELETE("users/:id", api.DeleteUser)
}
