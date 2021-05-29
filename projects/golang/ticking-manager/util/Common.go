package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func InitConfig() {
	viper.SetConfigName("Config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/") 
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

type Message struct {
	Code  int
	Err error
	Message string
	Data interface{}
}

func SendMessage(c *gin.Context, msg Message) {
	// err return
	if msg.Err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    msg.Code,
			"message": msg.Err.Error(),
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code":    msg.Code,
			"message": msg.Message,
			"data":    msg.Data,
		})
	}
}


func SendError(c *gin.Context, msg string) {
	// err return
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": msg,
	})
}
