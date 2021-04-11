package load_user_by_email

import (
	mocks2 "github.com/douglasdennys/go-mongodb/src/app/test"
	"github.com/douglasdennys/go-mongodb/src/domain/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApp_Add(t *testing.T) {
	f := test.MockAddUserParam()

	faker := mocks2.NewMockLoadUserRepositorySpy()
	app := NewDbLoadUserByEmail(faker)
	response, err := app.LoadByEmail(f.Email)
	assert.Nil(t, err)
	assert.Nil(t, response)

	fakers := mocks2.NewMockLoadUserRepositorySuccessSpy()
	apps := NewDbLoadUserByEmail(fakers)
	responses, errs := apps.LoadByEmail(f.Email)
	assert.Nil(t, errs)
	assert.NotNil(t, responses)
}

