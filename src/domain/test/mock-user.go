package test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/douglasdennys/go-mongodb/src/domain/entities"
	"github.com/douglasdennys/go-mongodb/src/domain/usecases/user"
)

func MockAddUserParam() *user.AddUserParam {
	mock := user.AddUserParam{}
	err := faker.FakeData(&mock)
	if err != nil {
		panic(err)
	}
	return &mock
}

func MockAddUserParamNotValid() *user.AddUserParam {
	mock := user.AddUserParam{}
	err := faker.FakeData(&mock)
	if err != nil {
		panic(err)
	}
	mock.Name = ""
	return &mock
}

func MockUserEntity() *entities.UserEntity {
	mock := entities.UserEntity{}
	err := faker.FakeData(&mock)
	if err != nil {
		panic(err)
	}
	return &mock
}
