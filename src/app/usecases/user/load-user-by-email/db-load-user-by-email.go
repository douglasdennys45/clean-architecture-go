package load_user_by_email

import (
	apps "github.com/douglasdennys/go-mongodb/src/app/protocols/db/user"
	"github.com/douglasdennys/go-mongodb/src/domain/entities"
	"github.com/douglasdennys/go-mongodb/src/domain/usecases/user"
)

type dbLoadUserByEmail interface {
	user.LoadUserByEmail
}

type app struct {
	load apps.LoadUserByEmailRepository
}

func NewDbLoadUserByEmail(load apps.LoadUserByEmailRepository) dbLoadUserByEmail {
	return &app{load}
}

func (a app) LoadByEmail(email string) (*entities.UserEntity, error) {
	response, err := a.load.LoadByEmail(email)
	if response != nil {
		return &response.UserEntity, err
	}
	return nil, err
}
