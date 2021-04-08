package mocks

import (
	"github.com/bxcodec/faker/v3"
	"github.com/douglasdennys/go-mongodb/src/app/protocols"
)

func MockAddUserParamRepo() *protocols.AddUserParamRepo {
	mock := protocols.AddUserParamRepo{}
	err := faker.FakeData(&mock)
	if err != nil {
		panic(err)
	}
	return &mock
}