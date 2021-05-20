package src 

import (
	"github.com/gin-gonic/gin"
	"fmt"
)


func Home(c *gin.Context) {
	fmt.Println("home")
}


func GetVersion(c *gin.Context) {
	c.JSON(200, gin.H{"version": "1.0", "app": "ticking-manager"})
}