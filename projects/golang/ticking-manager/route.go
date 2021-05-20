package route

import (
	"github.com/gin-gonic/gin"
)

func DefinitionRoute(router *gin.Engine) {
	router.GET("/ping", ping)
}


func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong", })
}
