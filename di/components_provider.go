package di

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDbClient() (*mongo.Client, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:7017")
	mongoClient, error := mongo.Connect(ctx, clientOptions)

	if error != nil {
		return nil, nil, errors.New("Error trying to connect to mongo database")
	}

	return mongoClient, cancel, nil
}
