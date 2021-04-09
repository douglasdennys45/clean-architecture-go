package test

import (
	"github.com/douglasdennys/go-mongodb/src/app/protocols/db/user"
)

/** Created mock interface AddUserRepository */
type addUserRepositorySpy struct {}

func NewMockAddUserRepositorySpy() user.AddUserRepository {
	return &addUserRepositorySpy{}
}

func (mock *addUserRepositorySpy) Add(addUser *user.AddUserParamRepo) error {
	return nil
}
/** Created mock interface AddUserRepository */