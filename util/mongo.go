package util

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect_Mongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))
	collection := client.Database("testing").Collection("numbers")
	collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.1419}})
}
