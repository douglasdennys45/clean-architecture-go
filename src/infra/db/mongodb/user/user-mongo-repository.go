package user

import (
	"context"
	"github.com/douglasdennys/go-mongodb/src/app/protocols/db/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository interface {
	user.AddUserRepository
}

type repository struct {
	collection *mongo.Collection
}

func NewUserMongoRepository(collection *mongo.Collection) UserMongoRepository {
	return &repository{collection}
}

func (r repository) Add(addUser *user.AddUserParamRepo) error {
	_, err := r.collection.InsertOne(context.Background(), addUser)
	return err
}
