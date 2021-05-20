package route

import (
	"github.com/gin-gonic/gin"
	"ticking-manager/src"
)

func DefinitionRoute(router *gin.Engine) {
	router.GET("/", src.Home)
	router.GET("/api/version", src.GetVersion)
}

