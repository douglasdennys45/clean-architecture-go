package user

import (
	"context"
	"github.com/douglasdennys/go-mongodb/src/app/protocols"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository interface {
	protocols.AddUserRepository
}

type repository struct {
	collection *mongo.Collection
}

func NewUserMongoRepository(collection *mongo.Collection) UserMongoRepository {
	return &repository{collection}
}

func (r repository) Add(addUser *protocols.AddUserParamRepo) error {
	_, err := r.collection.InsertOne(context.Background(), addUser)
	return err
}
