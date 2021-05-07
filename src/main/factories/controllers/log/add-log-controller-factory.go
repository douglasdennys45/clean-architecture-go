package log

import (
	"github.com/douglasdennys/clean-architecture-go/src/adapters/controllers/usecases"
	"github.com/douglasdennys/clean-architecture-go/src/adapters/controllers/usecases/log"
	"github.com/douglasdennys/clean-architecture-go/src/main/factories/usecases/repositories"
)

func MakeAddLogControllerFactory() usecases.Controller {
	return log.NewAddLogController(repositories.MakeDbAddLogRepositoryFactory())
}
