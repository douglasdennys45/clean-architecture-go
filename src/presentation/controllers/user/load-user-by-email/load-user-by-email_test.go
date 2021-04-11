package load_user_by_email

import (
	"encoding/json"
	"github.com/douglasdennys/go-mongodb/src/presentation/protocols"
	presentation "github.com/douglasdennys/go-mongodb/src/presentation/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

type searchs struct {
	Email string `json:"email"`
}

func TestController_Handle(t *testing.T) {
	var httpRequest protocols.HttpRequest
	var body searchs
	body.Email = "test@test.com"
	cur, _ := json.Marshal(body)
	httpRequest.Body = cur
	loadUserByEmailSpy := presentation.NewMockLoadUserByEmailSpy()
	controller := NewDbLoadUserByEmailController(loadUserByEmailSpy)
	response := controller.Handle(httpRequest)
	assert.Equal(t, 200, response.Code)

	loadUserByEmailErrorSpy := presentation.NewMockLoadUserByEmailErrorSpy()
	controllers := NewDbLoadUserByEmailController(loadUserByEmailErrorSpy)
	responses := controllers.Handle(httpRequest)
	assert.Equal(t, 500, responses.Code)
}