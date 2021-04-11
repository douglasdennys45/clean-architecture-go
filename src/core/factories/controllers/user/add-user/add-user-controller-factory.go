package add_user

import (
	factories "github.com/douglasdennys/go-mongodb/src/core/factories/controllers"
	usecases "github.com/douglasdennys/go-mongodb/src/core/factories/usecases/user/add-user"
	controllers "github.com/douglasdennys/go-mongodb/src/presentation/controllers/user/add-user"
	"github.com/douglasdennys/go-mongodb/src/presentation/protocols"
)

func MakeAddUserController() protocols.Controller {
	return controllers.NewDbAddUserController(usecases.MakeDbAddUser(), factories.MakeValidation())
}