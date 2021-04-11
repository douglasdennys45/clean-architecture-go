package add_user

import (
	apps "github.com/douglasdennys/go-mongodb/src/app/usecases/user/load-user-by-email"
	"github.com/douglasdennys/go-mongodb/src/domain/usecases/user"
	"github.com/douglasdennys/go-mongodb/src/infra/db/mongodb/helpers"
	repos "github.com/douglasdennys/go-mongodb/src/infra/db/mongodb/user"
)

func MakeDbLoadUserByEmail() user.LoadUserByEmail {
	collection := helpers.GetCollection("users")
	repo := repos.NewUserMongoRepository(collection)

	return apps.NewDbLoadUserByEmail(repo)
}
