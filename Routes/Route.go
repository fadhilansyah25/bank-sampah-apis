package Routes

import (
	LoginRoute "golang-final-project/Routes/Login"
	UsersRoute "golang-final-project/Routes/Users"

	"github.com/labstack/echo/v4"
)

func RouteVersion1() *echo.Echo {
	e := echo.New()
	r1 := e.Group("v1/")
	UsersRoute.UsersRouter(r1)
	LoginRoute.UserLoginRoute(r1)

	return e
}
