package config

import (
	"github.com/douglasdennys/go-mongodb/src/main/routes"
	"github.com/labstack/echo/v4"
)

func Routes(router *echo.Echo) {
	routes.UserRouter(router)
}