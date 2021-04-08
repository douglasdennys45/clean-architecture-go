package add_user

import (
	"github.com/douglasdennys/go-mongodb/src/app/protocols"
	"github.com/douglasdennys/go-mongodb/src/domain/usecases/user"
)

type DbAddUser interface {
	user.AddUser
}

type app struct {
	addUserRepo protocols.AddUserRepository
}

func NewDbAddUser(addUserRepo protocols.AddUserRepository) DbAddUser {
	return &app{addUserRepo}
}

func (a *app) Add(addUser *user.AddUserParam) error {
	param := protocols.AddUserParamRepo{addUser}
	err := a.addUserRepo.Add(&param)
	return err
}
