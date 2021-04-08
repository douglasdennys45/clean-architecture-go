package mocks

import (
	"github.com/douglasdennys/go-mongodb/src/app/protocols"
)

/** Created mock interface AddUserRepository */
type addUserRepositorySpy struct {}

func NewMockAddUserRepositorySpy() protocols.AddUserRepository {
	return &addUserRepositorySpy{}
}

func (mock *addUserRepositorySpy) Add(addUser *protocols.AddUserParamRepo) error {
	return nil
}
/** Created mock interface AddUserRepository */