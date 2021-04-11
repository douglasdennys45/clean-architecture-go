package user

import "github.com/douglasdennys/go-mongodb/src/domain/entities"

type LoadUserByEmail interface {
	LoadByEmail(email string) (*entities.UserEntity, error)
}