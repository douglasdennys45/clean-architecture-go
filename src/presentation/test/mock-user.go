package test

import (
	"errors"
	"github.com/douglasdennys/go-mongodb/src/domain/entities"
	"github.com/douglasdennys/go-mongodb/src/domain/usecases/user"
)

type addUserSpy struct {}

func NewMockAddUserSpy() user.AddUser {
	return &addUserSpy{}
}

func (mock *addUserSpy) Add(addUser *user.AddUserParam) error {
	return nil
}

type addUserNotValidSpy struct {}

func NewMockAddUserNotValidSpy() user.AddUser {
	return &addUserNotValidSpy{}
}

func (mock *addUserNotValidSpy) Add(addUser *user.AddUserParam) error {
	return errors.New("error")
}

type loadUserByEmailSpy struct {}

func NewMockLoadUserByEmailSpy() user.LoadUserByEmail {
	return &loadUserByEmailSpy{}
}

func (mock *loadUserByEmailSpy) LoadByEmail(email string) (*entities.UserEntity, error) {
	return nil, nil
}

type loadUserByEmailErrorSpy struct {}

func NewMockLoadUserByEmailErrorSpy() user.LoadUserByEmail {
	return &loadUserByEmailErrorSpy{}
}

func (mock *loadUserByEmailErrorSpy) LoadByEmail(email string) (*entities.UserEntity, error) {
	return nil, errors.New("error")
}