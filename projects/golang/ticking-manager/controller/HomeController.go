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
	var uc *UserController
	var mc *MovieController
	users, errUsers := uc.getUserController().Service.FindAllUsers()
	if errUsers != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": errUsers,
		})
	}

	movies, errMovies := mc.getMovieController().Service.FindAllMovies()
	if errMovies != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": errMovies,
		})
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "票务管理系统后台",
			"users": users,
			"movies": movies,
	})
}


func (hc *HomeController) AddUser(c *gin.Context) {

	c.HTML(http.StatusOK, "add.tmpl", gin.H{
			"title": "添加用户",
		})
}


func (hc *HomeController) AddMovie(c *gin.Context) {

	c.HTML(http.StatusOK, "addmovie.tmpl", gin.H{
		"title": "添加电影",
	})
}













