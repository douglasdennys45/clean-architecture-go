package add_user

import (
	user2 "github.com/douglasdennys/go-mongodb/src/app/protocols/db/user"
	"github.com/douglasdennys/go-mongodb/src/domain/usecases/user"
	bcrypt_adapter "github.com/douglasdennys/go-mongodb/src/infra/criptography/bcrypt-adapter"
)

type DbAddUser interface {
	user.AddUser
}

type app struct {
	addUserRepo user2.AddUserRepository
	crypto bcrypt_adapter.BcryptAdapter
}

func NewDbAddUser(addUserRepo user2.AddUserRepository, crypto bcrypt_adapter.BcryptAdapter) DbAddUser {
	return &app{addUserRepo, crypto}
}

func (a *app) Add(addUser *user.AddUserParam) error {
	password, err := a.crypto.Hash(addUser.Password)
	param := user2.AddUserParamRepo{addUser}
	param.Password = password
	err = a.addUserRepo.Add(&param)
	return err
}
