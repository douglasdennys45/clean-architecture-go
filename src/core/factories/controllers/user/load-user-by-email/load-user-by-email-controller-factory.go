package add_user

import (
	usecases "github.com/douglasdennys/go-mongodb/src/core/factories/usecases/user/load-user-by-email"
	controllers "github.com/douglasdennys/go-mongodb/src/presentation/controllers/user/load-user-by-email"
	"github.com/douglasdennys/go-mongodb/src/presentation/protocols"
)

func MakeLoadUserByEmailController() protocols.Controller {
	return controllers.NewDbLoadUserByEmailController(usecases.MakeDbLoadUserByEmail())
}