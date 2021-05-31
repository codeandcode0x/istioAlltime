package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket-manager/controller"
	"ticket-manager/middleware"
)

func DefinitionRoute(router *gin.Engine) {
	// middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.UseCookieSession())
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

	auth := router.Group("/")
	auth.Use(middleware.AuthMiddle())
	{
		auth.GET("/", homeController.ProxyHome)
		auth.GET("/logout", homeController.Logout)
		auth.GET("/user/add", homeController.AddUser)
		auth.GET("/movie/add", homeController.AddMovie)

		// user
		var userController *controller.UserController
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
	api.GET("/movie/:id", movieController.GetMovieByID)

}

// no route
func NoRouteResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code": 404,
		"error":  "oops, page not exists!",
	})
}


