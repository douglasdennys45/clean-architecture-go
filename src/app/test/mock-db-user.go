package test

import (
	"github.com/douglasdennys/go-mongodb/src/app/protocols/db/user"
	"github.com/douglasdennys/go-mongodb/src/domain/test"
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
type loadUserRepositorySpy struct {}

func NewMockLoadUserRepositorySpy() user.LoadUserByEmailRepository {
	return &loadUserRepositorySpy{}
}

func (mock *loadUserRepositorySpy) LoadByEmail(email string) (*user.UserRepo, error) {
	return nil, nil
}

type loadUserRepositorySuccessSpy struct {}

func NewMockLoadUserRepositorySuccessSpy() user.LoadUserByEmailRepository {
	return &loadUserRepositorySuccessSpy{}
}

func (mock *loadUserRepositorySuccessSpy) LoadByEmail(email string) (*user.UserRepo, error) {
	var userRepo user.UserRepo
	userRepo.UserEntity = test.MockUserEntity()
	return &userRepo, nil
}