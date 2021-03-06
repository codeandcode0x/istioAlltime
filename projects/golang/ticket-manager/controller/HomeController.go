package controller

import (
	"net"
	"net/http"
	"ticket-manager/middleware"
	"ticket-manager/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type HomeController struct {
	apiVersion string
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
			"code":  -1,
			"error": "user not find or find err !",
		})
	} else {
		errPasswd := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if user.Email == email && errPasswd == nil {
			middleware.SaveAuthSession(c, user.ID)
			hc.ProxyHome(c)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":  -1,
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
	pagesize := viper.GetInt("PAGE_SIZE")
	users, errUsers := us.FindAllUsers()
	if errUsers != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": errUsers,
		})
	}

	movies, errMovies := ms.FindMovieByPages(1, pagesize)
	if errMovies != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": errMovies,
		})
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":  "票务管理系统后台",
		"users":  users,
		"movies": movies,
	})
}

func (hc *HomeController) UserList(c *gin.Context) {
	var us *service.UserService
	users, errUsers := us.FindAllUsers()
	if errUsers != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": errUsers,
		})
	}

	c.HTML(http.StatusOK, "users.tmpl", gin.H{
		"title": "用户列表",
		"users": users,
	})
}

func (hc *HomeController) AddUser(c *gin.Context) {
	c.HTML(http.StatusOK, "add.tmpl", gin.H{
		"title": "添加用户",
	})
}

func (hc *HomeController) MovieList(c *gin.Context) {
	var ms *service.MovieService
	pagesize := viper.GetInt("PAGE_SIZE")
	movies, errMovies := ms.FindMovieByPages(1, pagesize)
	if errMovies != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": errMovies,
		})
	}

	c.HTML(http.StatusOK, "movies.tmpl", gin.H{
		"title":  "电影列表",
		"movies": movies,
	})
}

func (hc *HomeController) AddMovie(c *gin.Context) {
	c.HTML(http.StatusOK, "addmovie.tmpl", gin.H{
		"title": "添加电影",
	})
}

func (hc *HomeController) ShowList(c *gin.Context) {
	var ms *service.ShowService
	pagesize := viper.GetInt("PAGE_SIZE")
	shows, errShows := ms.FindShowByPagesWithKeys("演唱会", 1, pagesize)
	if errShows != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": errShows,
		})
	}

	c.HTML(http.StatusOK, "shows.tmpl", gin.H{
		"title": "添加演出",
		"shows": shows,
	})
}

func (hc *HomeController) AddShow(c *gin.Context) {
	c.HTML(http.StatusOK, "addshow.tmpl", gin.H{
		"title": "添加演出",
	})
}

func (hc *HomeController) InfoList(c *gin.Context) {
	var ms *service.InfoService
	pagesize := viper.GetInt("PAGE_SIZE")
	infos, errInfo := ms.FindInfoByPages(1, pagesize)
	if errInfo != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": errInfo,
		})
	}

	c.HTML(http.StatusOK, "infos.tmpl", gin.H{
		"title": "添加资讯",
		"infos": infos,
	})
}

func (hc *HomeController) AddInfo(c *gin.Context) {
	c.HTML(http.StatusOK, "addinfo.tmpl", gin.H{
		"title": "添加资讯",
	})
}

func (hc *HomeController) Healthz(c *gin.Context) {
	conn, err := net.DialTimeout("tcp", "http://mariadb:3306", time.Second*3)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	err = conn.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "health",
	})
}
