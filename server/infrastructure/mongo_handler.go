package infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoHandler struct {
	client *mongo.Client
}

func NewMongoHandler() (*MongoHandler, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	mongoHandler := MongoHandler{
		client: client,
	}
	return &mongoHandler, nil
}
