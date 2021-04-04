package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var StringURL string
var DbName string
var Conn *mongo.Client
var Db *mongo.Database

func Connect(url string, name string) error {
	StringURL = url
	DbName = name
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return err
	}
	Conn = client
	Db = client.Database(name)
	return nil
}

func GetCollection(name string) *mongo.Collection {
	return Db.Collection(name)
}

func Disconnect() error {
	if err := Conn.Disconnect(context.Background()); err != nil {
		panic(err)
	}
	return nil
}
