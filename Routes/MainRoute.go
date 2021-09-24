package Routes

import (
	"golang-final-project/Middleware"
	BankSampahRoute "golang-final-project/Routes/BankSampah"
	JenisSampahRoute "golang-final-project/Routes/JenisSampah"
	LoginRoute "golang-final-project/Routes/Login"
	OperatorSampahRoute "golang-final-project/Routes/OperatorSampah"
	TransactionRoute "golang-final-project/Routes/Transaction"
	UsersRoute "golang-final-project/Routes/Users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteVersion1() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} => ${host} ${method} uri=${uri}  ||  ${status} => ${latency_human} ${remote_ip} \n",
	}))
	e.Use(middleware.BodyDump(Middleware.Log))

	r1 := e.Group("api/v1/")
	UsersRoute.UsersRouter(r1)
	LoginRoute.UserLoginRoute(r1)
	BankSampahRoute.BankSampahRouter(r1)
	OperatorSampahRoute.OperatorSampahRoute(r1)
	JenisSampahRoute.JenisSampahRouter(r1)
	TransactionRoute.TransactionRouter(r1)

	return e
}
