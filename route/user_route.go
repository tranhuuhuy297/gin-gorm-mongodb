package route

import (
	"gin-mongodb/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controller.CreateUser())
}
