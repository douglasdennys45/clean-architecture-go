package config

import (
	"github.com/labstack/echo/v4"
)

func StartApp() {
	router := echo.New()
	Routes(router)

	router.Logger.Fatal(router.Start(":8080"))
}