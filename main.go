package main

import (
	"gin-mongodb/util"

	"github.com/gin-gonic/gin"
)

func main() {
	util.Load_Env()
	util.Connect_Mongo()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})

	endpoint := "localhost:" + util.PORT
	router.Run(endpoint)
}
