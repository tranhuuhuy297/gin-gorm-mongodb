package util

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	MongoClient, _ = mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))
}

func GetCollection(collection_name string) *mongo.Collection {
	collection := MongoClient.Database(DATABASE).Collection(collection_name)

	return collection
}
