package load_user_by_email

import (
	"encoding/json"
	"github.com/douglasdennys/go-mongodb/src/domain/usecases/user"
	"github.com/douglasdennys/go-mongodb/src/presentation/helpers/http"
	"github.com/douglasdennys/go-mongodb/src/presentation/protocols"
)

type loadUserByEmailController interface {
	protocols.Controller
}

type controller struct {
	loadUser user.LoadUserByEmail
}

func NewDbLoadUserByEmailController(loadUser user.LoadUserByEmail) loadUserByEmailController {
	return &controller{loadUser}
}

type search struct {
	Email string `json:"email"`
}

func (c controller) Handle(httpRequest protocols.HttpRequest) protocols.HttpResponse {
	var body search
	_ = json.Unmarshal(httpRequest.Body, &body)

	response, err := c.loadUser.LoadByEmail(body.Email)
	if err != nil {
		return http.ServerError(err)
	}
	return http.Ok(response)
}
