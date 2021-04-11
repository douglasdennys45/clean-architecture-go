package routes

import (
	"github.com/douglasdennys/go-mongodb/src/core/adapters"
	controller "github.com/douglasdennys/go-mongodb/src/core/factories/controllers/user/add-user"
	controllers "github.com/douglasdennys/go-mongodb/src/core/factories/controllers/user/load-user-by-email"
	"github.com/labstack/echo/v4"
)

func UserRouter(router *echo.Echo) {
	router.POST("/v1/users", adapters.AdapterRouter(controller.MakeAddUserController()))
	router.GET("/v1/users", adapters.AdapterRouter(controllers.MakeLoadUserByEmailController()))
}