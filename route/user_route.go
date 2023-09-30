package route

import (
	"gin-mongodb/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controller.CreateUser())
	router.GET("/user/:userId", controller.GetUser())
	router.PUT("/user/:userId", controller.UpdateUser())
	router.DELETE("/user/:userId", controller.DeleteUser())
	router.GET("/users", controller.GetUsers())
}
