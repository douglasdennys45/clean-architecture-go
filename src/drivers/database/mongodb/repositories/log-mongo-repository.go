package repositories

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AddLogRepository interface {
	Add(logData []byte) error
}

type Log struct {
	Error     string    `json:"name" validate:"required"`
	Service   string    `json:"service" validate:"required"`
	Uri       string    `json:"uri" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
}

type logMongoRepository interface {
	AddLogRepository
}

type repo struct {
	collection *mongo.Collection
}

func NewLogMongoRepository(collection *mongo.Collection) logMongoRepository {
	return &repo{collection}
}

func (r *repo) Add(logData []byte) error {
	var log Log
	_ = json.Unmarshal(logData, &log)
	_, err := r.collection.InsertOne(context.TODO(), log)
	return err
}