package user

import "github.com/douglasdennys/go-mongodb/src/domain/entities"

type UserRepo struct {
	entities.UserEntity
}

type LoadUserByEmailRepository interface {
	LoadByEmail(email string) (*UserRepo, error)
}
