package add_user

import (
	"errors"
	mocks2 "github.com/douglasdennys/go-mongodb/src/app/mocks"
	"github.com/douglasdennys/go-mongodb/src/app/protocols"
	"github.com/douglasdennys/go-mongodb/src/domain/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApp_Add(t *testing.T) {
	f := mocks.MockAddUserParam()
	mock := protocols.AddUserParamRepo{f}

	faker := mocks2.NewMockAddUserRepositorySpy()
	app := NewDbAddUser(faker)

	err := app.Add(mock.AddUserParam)
	assert.Nil(t, err)

	err2 := app.Add(mock.AddUserParam)
	err2 = errors.New("error")
	assert.NotNil(t, err2)
}