package main

import (
	"gin-mongodb/route"
	"gin-mongodb/util"

	"github.com/gin-gonic/gin"
)

func main() {
	util.LoadEnv()
	util.ConnectMongo()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})

	route.UserRoute(router)

	endpoint := "localhost:" + util.PORT
	router.Run(endpoint)
}
