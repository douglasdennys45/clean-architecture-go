package protocols

import "github.com/douglasdennys/go-mongodb/src/domain/usecases/user"

type AddUserParamRepo struct {
	*user.AddUserParam
}

type AddUserRepository interface {
	Add(addUser *AddUserParamRepo) error
}