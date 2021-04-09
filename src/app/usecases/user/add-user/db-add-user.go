package add_user

import (
	user2 "github.com/douglasdennys/go-mongodb/src/app/protocols/db/user"
	"github.com/douglasdennys/go-mongodb/src/domain/usecases/user"
)

type DbAddUser interface {
	user.AddUser
}

type app struct {
	addUserRepo user2.AddUserRepository
}

func NewDbAddUser(addUserRepo user2.AddUserRepository) DbAddUser {
	return &app{addUserRepo}
}

func (a *app) Add(addUser *user.AddUserParam) error {
	param := user2.AddUserParamRepo{addUser}
	err := a.addUserRepo.Add(&param)
	return err
}
