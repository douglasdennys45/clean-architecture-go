package user

import (
	"errors"
	"github.com/benweissmann/memongo"
	"github.com/douglasdennys/go-mongodb/src/app/protocols"
	"github.com/douglasdennys/go-mongodb/src/domain/test"
	"github.com/douglasdennys/go-mongodb/src/infra/db/mongodb/helpers"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"testing"
)

func NewMockUserMongoRepositorySpy() *mongo.Collection {
	mongoServer, err := memongo.Start("4.0.5")
	if err != nil {
		log.Fatal(err)
	}

	_ = helpers.Connect(mongoServer.URI(), memongo.RandomDatabase())
	collection := helpers.GetCollection("users")
	return collection
}

func TestRepository_Add(t *testing.T) {
	f := test.MockAddUserParam()
	mock := protocols.AddUserParamRepo{f}

	faker := NewMockUserMongoRepositorySpy()
	repo := NewUserMongoRepository(faker)

	err := repo.Add(&mock)
	assert.Nil(t, err)

	err = repo.Add(&mock)
	err = errors.New("error")
	assert.NotNil(t, err)
}
