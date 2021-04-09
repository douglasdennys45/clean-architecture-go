package add_user

import (
	apps "github.com/douglasdennys/go-mongodb/src/app/usecases/user/add-user"
	"github.com/douglasdennys/go-mongodb/src/domain/usecases/user"
	bcrypt_adapter "github.com/douglasdennys/go-mongodb/src/infra/criptography/bcrypt-adapter"
	"github.com/douglasdennys/go-mongodb/src/infra/db/mongodb/helpers"
	repos "github.com/douglasdennys/go-mongodb/src/infra/db/mongodb/user"
)

func MakeDbAddUser() user.AddUser {
	collection := helpers.GetCollection("users")
	repo := repos.NewUserMongoRepository(collection)
	crypto := bcrypt_adapter.NewBcryptAdapter()

	return apps.NewDbAddUser(repo, crypto)
}
