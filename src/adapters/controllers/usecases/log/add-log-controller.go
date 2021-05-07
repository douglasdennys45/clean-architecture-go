package log

import (
	"github.com/douglasdennys/clean-architecture-go/src/adapters/controllers/ports"
	"github.com/douglasdennys/clean-architecture-go/src/adapters/controllers/shared"
	"github.com/douglasdennys/clean-architecture-go/src/adapters/controllers/usecases"
	"github.com/douglasdennys/clean-architecture-go/src/application/ports/repositories/log"
)

type addLogController interface {
	usecases.Controller
}

type controller struct {
	add log.AddLogRepository
}

func NewAddLogController(add log.AddLogRepository) addLogController {
	return &controller{add}
}

func (c *controller) Handle(req ports.HttpRequest) ports.HttpResponse {
	err := c.add.Add(req.Body)
	if err != nil {
		return shared.ServerError(err.Error())
	}
	return shared.NoContent()
}
