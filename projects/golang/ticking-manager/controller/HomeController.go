package controller 

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"ticket-manager/middleware"
	"ticket-manager/service"
	// "fmt"
)

type HomeController struct {
	apiVersion  string
}


// get controller
func (hc *HomeController) getHomeController() *HomeController {
	return &HomeController{"v1"}
}


func (hc *HomeController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "用户登录",
	})
}


func (hc *HomeController) DoLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	var us *service.UserService
	user, errFind := us.FindUserByEmail(email)

	if user == nil || errFind != nil {
		
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": "user not find or find err !",
		})
	}else {
		errPasswd := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if user.Email == email && errPasswd == nil {
			middleware.SaveAuthSession(c, user.ID)
			hc.ProxyHome(c)
		}else{
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"error": "email or password error !",
			})
		}
	}
}


func (hc *HomeController) Logout(c *gin.Context) {
	middleware.ClearAuthSession(c)
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "用户登录",
	})
}


func (hc *HomeController) ProxyHome(c *gin.Context) {
	var us *service.UserService
	var ms *service.MovieService
	users, errUsers := us.FindAllUsers()
	if errUsers != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": errUsers,
		})
	}

	movies, errMovies := ms.FindAllMovies()
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
















