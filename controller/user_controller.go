package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("post rest")
	}
}
