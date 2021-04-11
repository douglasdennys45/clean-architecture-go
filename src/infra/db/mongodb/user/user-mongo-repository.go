package user

import (
	"context"
	"github.com/douglasdennys/go-mongodb/src/app/protocols/db/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserMongoRepository interface {
	user.AddUserRepository
	user.LoadUserByEmailRepository
}

type userData struct {
	Name      string    `json:"name" bson:"name" validate:"required"`
	Email     string    `json:"email" bson:"email" validate:"required,email"`
	Password  string    `json:"password" bson:"password" validate:"required"`
	Active    bool      `json:"active" bson:"active" validate:"required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" validate:"datetime"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at" validate:"datetime"`
}

type repository struct {
	collection *mongo.Collection
}

func NewUserMongoRepository(collection *mongo.Collection) UserMongoRepository {
	return &repository{collection}
}

func (r *repository) Add(addUser *user.AddUserParamRepo) error {
	_, err := r.collection.InsertOne(context.Background(), &userData{addUser.AddUserParam.Name, addUser.AddUserParam.Email, addUser.AddUserParam.Password, true, time.Now(), time.Now()})
	return err
}

func (r *repository) LoadByEmail(email string) (*user.UserRepo, error) {
	var userRepo user.UserRepo
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&userRepo.UserEntity)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return nil, nil
	}
	return &userRepo, err
}
