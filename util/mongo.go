package util

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))

	return client
}

var DB *mongo.Client = ConnectMongo()

func GetCollection(collection_name string) *mongo.Collection {
	collection := DB.Database(DATABASE).Collection(collection_name)

	return collection
}
