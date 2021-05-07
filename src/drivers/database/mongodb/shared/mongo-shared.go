package shared

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	URI string
	DB_NAME string
	CLIENT *mongo.Client
	DB *mongo.Database
)

func Connect(uri string, dbName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	DB = client.Database(dbName)
	CLIENT = client
	DB_NAME = dbName
	URI = uri
	return nil
}

func GetCollection(col string) *mongo.Collection {
	return DB.Collection(col)
}