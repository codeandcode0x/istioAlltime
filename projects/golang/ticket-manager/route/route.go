package route

import (
	"ticket-manager/controller"
	"ticket-manager/middleware"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	TimeDuration = 5
)

func DefinitionRoute(router *gin.Engine) {
	// middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.UseCookieSession())
	router.Use(middleware.TimeoutHandler(time.Second * TimeDuration))
	// no route
	router.NoRoute(NoRouteResponse)
	// home
	var homeController *controller.HomeController

	router.Static("/web/assets", "./web/assets")
	router.StaticFS("/web/upload", http.Dir("/web/upload"))
	router.LoadHTMLGlob("web/*.tmpl")

	router.GET("/login", homeController.Login)
	router.POST("/dologin", homeController.DoLogin)

	var movieController *controller.MovieController
	var userController *controller.UserController

	auth := router.Group("/")
	auth.Use(middleware.AuthMiddle())
	{
		auth.GET("/", homeController.ProxyHome)
		auth.GET("/logout", homeController.Logout)
		auth.GET("/user/add", homeController.AddUser)
		auth.GET("/movie/add", homeController.AddMovie)

		// user
		auth.GET("/users", userController.GetAllUsers)
		auth.POST("/user/create", userController.CreateUser)
		auth.POST("/user/update", userController.UpdateUser)
		auth.POST("/user/delete", userController.DeleteUser)

		// movie
		auth.GET("/movies", movieController.GetAllMovies)
		auth.POST("/movie/create", movieController.CreateMovie)
		auth.POST("/movie/update", movieController.UpdateMovie)
		auth.POST("/movie/delete", movieController.DeleteMovie)
	}

	// api
	api := router.Group("/api")
	api.GET("/movies", movieController.GetAllMovies)
	api.GET("/rpc/movies", movieController.GetAllMoviesRPC)
	api.GET("/movie/:id", movieController.GetMovieByID)
	api.GET("/users", userController.GetAllUsers)

}

// no route
func NoRouteResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":  404,
		"error": "oops, page not exists!",
	})
}
