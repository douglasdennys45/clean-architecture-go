package add_user

import (
	"encoding/json"
	add_user "github.com/douglasdennys/go-mongodb/src/app/usecases/user/add-user"
	"github.com/douglasdennys/go-mongodb/src/domain/usecases/user"
	"github.com/douglasdennys/go-mongodb/src/presentation/helpers/http"
	"github.com/douglasdennys/go-mongodb/src/presentation/protocols"
)

type AddUserController interface {
	protocols.Controller
}

type controller struct {
	addUser user.AddUser
	valid protocols.Validation
}

func NewAddUserController(add add_user.DbAddUser, valid protocols.Validation) AddUserController {
	return &controller{add, valid}
}

func (c controller) Handle(httpRequest protocols.HttpRequest) protocols.HttpResponse {
	var body user.AddUserParam
	_ = json.Unmarshal(httpRequest.Body, &body)

	err := c.valid.Validate(body)
	if err != nil {
		return http.BadRequest(err)
	}

	err = c.addUser.Add(&body)
	if err != nil {
		return http.ServerError(err)
	}
	return http.Created(nil)
}
