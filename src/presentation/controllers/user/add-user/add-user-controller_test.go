package add_user

import (
	"encoding/json"
	"github.com/douglasdennys/go-mongodb/src/domain/test"
	"github.com/douglasdennys/go-mongodb/src/presentation/protocols"
	presentation "github.com/douglasdennys/go-mongodb/src/presentation/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestController_Handle(t *testing.T) {
	var httpRequest protocols.HttpRequest
	cur, _ := json.Marshal(test.MockAddUserParam())
	httpRequest.Body = cur
	addUserSpy := presentation.NewMockAddUserSpy()
	validationSpy := presentation.NewMockValidateSpy()
	controller := NewAddUserController(addUserSpy, validationSpy)
	response := controller.Handle(httpRequest)
	assert.Equal(t, 201, response.Code)
	assert.Equal(t, nil, response.Data)
}

func TestController_HandleNotValid(t *testing.T) {
	var httpRequest protocols.HttpRequest
	cur, _ := json.Marshal(test.MockAddUserParamNotValid())
	httpRequest.Body = cur
	addUserSpy := presentation.NewMockAddUserSpy()
	validationSpy := presentation.NewMockValidateNotValidSpy()
	controller := NewAddUserController(addUserSpy, validationSpy)
	response := controller.Handle(httpRequest)
	assert.Equal(t, 400, response.Code)
}

func TestController_HandleError(t *testing.T) {
	var httpRequest protocols.HttpRequest
	cur, _ := json.Marshal(test.MockAddUserParam())
	httpRequest.Body = cur
	addUserSpy := presentation.NewMockAddUserNotValidSpy()
	validationSpy := presentation.NewMockValidateSpy()
	controller := NewAddUserController(addUserSpy, validationSpy)
	response := controller.Handle(httpRequest)
	assert.Equal(t, 500, response.Code)
}