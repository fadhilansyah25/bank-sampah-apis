package Routes

import (
	"golang-final-project/Controllers"

	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	r1 := e.Group("v1/")
	r1.POST("user", Controllers.UserRegister)

	return e
}
