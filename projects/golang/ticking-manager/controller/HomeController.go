package controller 

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
)

type HomeController struct {
	apiVersion  string
}


// get controller
func (hc *HomeController) getHomeController() *HomeController {
	return &HomeController{"v1"}
}


func (hc *HomeController) ProxyHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "票务管理系统后台",
		})
}