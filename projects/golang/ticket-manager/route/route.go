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
	router.Use(middleware.Tracing())
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
	var showController *controller.ShowController
	var infoController *controller.InfoController
	var orderController *controller.OrderController

	auth := router.Group("/")
	auth.Use(middleware.AuthMiddle())
	{
		auth.GET("/", homeController.ProxyHome)
		auth.GET("/logout", homeController.Logout)

		// user
		auth.GET("/users", homeController.UserList)   //web ui
		auth.GET("/user/add", homeController.AddUser) //web ui
		auth.POST("/user/create", userController.CreateUser)
		auth.POST("/user/update", userController.UpdateUser)
		auth.POST("/user/delete", userController.DeleteUser)

		// movie
		auth.GET("/movies", homeController.MovieList)   //web ui
		auth.GET("/movie/add", homeController.AddMovie) //web ui
		auth.POST("/movie/create", movieController.CreateMovie)
		auth.POST("/movie/update", movieController.UpdateMovie)
		auth.POST("/movie/delete", movieController.DeleteMovie)

		auth.GET("/shows", homeController.ShowList)   //web ui
		auth.GET("/show/add", homeController.AddShow) //web ui
		auth.POST("/show/create", showController.CreateShow)
		auth.POST("/show/update", showController.UpdateShow)
		auth.POST("/show/delete", showController.DeleteShow)

		auth.GET("/infos", homeController.InfoList)   //web ui
		auth.GET("/info/add", homeController.AddInfo) //web ui
		auth.POST("/info/create", infoController.CreateInfo)
		auth.POST("/info/update", infoController.UpdateInfo)
		auth.POST("/info/delete", infoController.DeleteInfo)
	}

	// api
	api := router.Group("/api")
	api.GET("/movies", movieController.GetAllMovies)
	api.GET("/rpc/movies", movieController.GetAllMoviesRPC)
	api.GET("/movie/:id", movieController.GetMovieByID)
	api.GET("/users", userController.GetAllUsers)
	api.GET("/shows", showController.GetAllShows)
	api.GET("/infos", infoController.GetAllInfos)

	api.GET("/orders", orderController.GetOrderByPages)
	api.POST("/order/create", orderController.CreateOrder)
	// api.POST("/order/create", orderController.GetOrderByPages)

	api.GET("/healthz", homeController.Healthz)
}

// no route
func NoRouteResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":  404,
		"error": "oops, page not exists!",
	})
}
