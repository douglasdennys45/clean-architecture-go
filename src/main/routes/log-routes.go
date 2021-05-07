package routes

import (
	"github.com/douglasdennys/clean-architecture-go/src/main/adapters"
	"github.com/douglasdennys/clean-architecture-go/src/main/factories/controllers/log"
	"net/http"
)

func LogRoutes() {
	http.HandleFunc("/v1/logs", adapters.AdaptRouter(log.MakeAddLogControllerFactory()))
}