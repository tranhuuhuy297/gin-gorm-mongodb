package controller

import (
	"context"
	"encoding/json"
	"gin-mongodb/model"
	"gin-mongodb/response"
	"gin-mongodb/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		var user_collection *mongo.Collection = util.GetCollection(util.USER_COLLECTION)
		var user model.User

		newUser := model.User{
			Name:       user.Name,
			Gmail:      user.Gmail,
			IsActivate: user.IsActivate,
		}

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, response.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		result, err := user_collection.InsertOne(ctx, newUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, response.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, response.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		var user_collection *mongo.Collection = util.GetCollection(util.USER_COLLECTION)

		userId := c.Param("userId")
		objId, _ := primitive.ObjectIDFromHex(userId)
		var user model.User
		err := user_collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// test redis
		user_obj, err := json.Marshal(user)
		util.RedisSet(userId, user_obj, 10)

		c.JSON(http.StatusOK, response.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
	}
}
