package infrastructure

import (
	"context"
	"reflect"

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
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		return nil, err
	}
	mongoHandler := MongoHandler{
		client:   client,
		database: client.Database("dchat"),
	}
	return &mongoHandler, nil
}

func (handler *MongoHandler) castArrayKvToD(doc []database.KV) bson.D {
	result := bson.D{}
	for _, kv := range doc {
		result = append(result, handler.castKvToE(kv))
	}
	return result
}

func (handler *MongoHandler) castKvToE(kv database.KV) primitive.E {
	if reflect.TypeOf(kv.Value).Elem() == reflect.TypeOf(database.KV{}) {
		kind := reflect.TypeOf(kv.Value).Kind()
		if kind == reflect.Array || kind == reflect.Slice {
			kv.Value = handler.castArrayKvToD(kv.Value.([]database.KV))
		} else {
			kv.Value = handler.castKvToE(kv.Value.(database.KV))
		}
	}
	e := primitive.E{
		Key:   kv.Key,
		Value: kv.Value,
	}
	return e
}

func (handler *MongoHandler) createIndexOptions(opts []database.KV) *options.IndexOptions {
	indexOptions := options.IndexOptions{}
	for _, opt := range opts {
		switch opt.Key {
		case "unique":
			indexOptions.Unique = opt.Value.(*bool)
		}
	}
	return &indexOptions
}

func (handler *MongoHandler) CreateIndex(collectionName string, index []database.KV, opt []database.KV) error {
	indexModel := mongo.IndexModel{
		Keys:    handler.castArrayKvToD(index),
		Options: handler.createIndexOptions(opt),
	}
	_, err := handler.database.Collection(collectionName).Indexes().CreateOne(context.Background(), indexModel)
	return err
}

func (handler *MongoHandler) Insert(collectionName string, doc []database.KV) (string, error) {
	result, err := handler.database.Collection(collectionName).InsertOne(context.Background(), handler.castArrayKvToD(doc))
	if err != nil {
		return "", err
	}
	return string(result.InsertedID.([]byte)), nil
}

func (handler *MongoHandler) Update(collectionName string, filter []database.KV, update []database.KV) error {
	_, err := handler.database.Collection(collectionName).UpdateMany(context.Background(), handler.castArrayKvToD(filter), handler.castArrayKvToD(update))
	return err
}
