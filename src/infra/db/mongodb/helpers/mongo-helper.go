package helpers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var stringURL string
var dbName string
var conn *mongo.Client
var db *mongo.Database

func Connect(url string, name string) error {
	stringURL = url
	dbName = name
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	conn = client
	db = client.Database(name)
	return err
}

func GetCollection(name string) *mongo.Collection {
	return db.Collection(name)
}

func Disconnect() error {
	return conn.Disconnect(context.Background())
}
