package protocols

import "github.com/douglasdennys/go-mongodb/src/domain/usecases/user"

type AddUserRepository interface {
	Add(addUser *user.AddUserParam) error
}