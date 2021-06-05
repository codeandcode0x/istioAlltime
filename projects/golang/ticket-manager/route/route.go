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
		"code":  404,
		"error": "oops, page not exists!",
	})
}

func timedHandler(duration time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		// get the underlying request context
		ctx := c.Request.Context()

		// create the response data type to use as a channel type
		type responseData struct {
			status int
			body   map[string]interface{}
		}

		// create a done channel to tell the request it's done
		doneChan := make(chan responseData)

		// here you put the actual work needed for the request
		// and then send the doneChan with the status and body
		// to finish the request by writing the response
		go func() {
			time.Sleep(duration)
			doneChan <- responseData{
				status: 200,
				body:   gin.H{"hello": "world"},
			}
		}()

		// non-blocking select on two channels see if the request
		// times out or finishes
		select {

		// if the context is done it timed out or was cancelled
		// so don't return anything
		case <-ctx.Done():
			return

			// if the request finished then finish the request by
			// writing the response
		case res := <-doneChan:
			c.JSON(res.status, res.body)
		}
	}
}
