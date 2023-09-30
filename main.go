package main

import (
	"gin-mongodb/route"
	"gin-mongodb/util"

	"github.com/gin-gonic/gin"
)

func main() {
	// init
	util.LoadEnv()
	util.ConnectMongo()
	util.ConnectRedis()

	router := gin.Default()
	endpoint := "localhost:" + util.PORT

	// health check
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	// add route
	route.UserRoute(router)

	// run server
	router.Run(endpoint)
}
