package controller

import (
	"context"
	"fmt"
	"gin-mongodb/model"
	"gin-mongodb/util"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		var user_collection *mongo.Collection = util.GetCollection(util.USER_COLLECTION)
		newUser := model.User{
			Name:       "huy",
			Gmail:      "tranhuuhuy297@gmail.com",
			IsActivate: true,
		}
		user_collection.InsertOne(ctx, newUser)
		fmt.Println("post rest")
	}
}
