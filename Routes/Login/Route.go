package LoginRoute

import (
	"golang-final-project/Configs/Database"
	"golang-final-project/Controllers/UserLoginHandler"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserLoginRoute(route *echo.Group) {
	jwtSecretKey := os.Getenv("SECRET_JWT")
	jwt := middleware.JWT([]byte(jwtSecretKey))

	api := UserLoginHandler.APIEnv{DB: Database.DB}

	route.POST("create-login", api.CreateUserLogin)
	route.POST("login", api.Login)
	route.GET("user-login", api.GetAlluserLogin, jwt)
	route.GET("user-login/:id", api.GetUserLoginByID, jwt)
	route.PUT("user-login/:id", api.UpdateUserLogin, jwt)
}
