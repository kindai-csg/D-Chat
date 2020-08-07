package infrastructure

import (
	"context"

	"github.com/kindai-csg/D-Chat/interfaces/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoHandler struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoHandler() (*MongoHandler, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	mongoHandler := MongoHandler{
		client:   client,
		database: client.Database("dchat"),
	}
	return &mongoHandler, nil
}

func (handler *MongoHandler) castKvToBson(doc []database.KV) bson.D {
	result := bson.D{}
	for _, kv := range doc {
		e := primitive.E{
			Key:   kv.Key,
			Value: kv.Value,
		}
		result = append(result, e)
	}
	return result
}

func (handler *MongoHandler) InsertOne(collectionName string, doc []database.KV) (string, error) {
	result, err := handler.database.Collection(collectionName).InsertOne(context.Background(), handler.castKvToBson(doc))
	if err != nil {
		return nil, err
	}
	return string(result.InsertedID.(primitive.ObjectID)), nil
}
