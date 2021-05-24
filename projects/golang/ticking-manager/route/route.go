package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket-manager/controller"
)

func DefinitionRoute(router *gin.Engine) {
	// middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// route
	// home
	var homeController *controller.HomeController
	router.GET("/", homeController.ProxyHome)
	router.Static("/web/assets", "./web/assets")
	router.StaticFS("/web/upload", http.Dir("/web/upload"))
	router.LoadHTMLGlob("web/*.tmpl")

	router.GET("/login", homeController.Login)
	router.GET("/dologin", homeController.DoLogin)
	router.GET("/logout", homeController.Logout)

	router.GET("/user/add", homeController.AddUser)
	router.GET("/movie/add", homeController.AddMovie)

	// user
	var userController *controller.UserController
	router.GET("/users", userController.GetAllUsers)
	router.POST("/user/create", userController.CreateUser)
	router.POST("/user/update", userController.UpdateUser)
	router.POST("/user/delete", userController.DeleteUser)

	// movie
	var movieController *controller.MovieController
	router.GET("/movies", movieController.GetAllMovies)
	router.POST("/movie/create", movieController.CreateMovie)
	router.POST("/movie/update", movieController.UpdateMovie)
	router.POST("/movie/delete", movieController.DeleteMovie)


	// authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
	// 	"admin": "admin", 
	// 	"demo":  "demo", 
	// }))

	// authorized.Use(AuthRequired())
	// {
	// 	 authorized.GET("/user", func(c *gin.Context) {
	//         // 获取用户，它是由 BasicAuth 中间件设置的
	//          c.JSON(http.StatusOK, gin.H{"OK": "OK"})
	//     })

	//     authorized.GET("/logout", func(c *gin.Context) {
	//         // 获取用户，它是由 BasicAuth 中间件设置的
	//          c.JSON(http.StatusOK, gin.H{"OK": "OK"})
	//     })
	// }

	// authorized.GET("/secrets", func(c *gin.Context) {
 //        // 获取用户，它是由 BasicAuth 中间件设置的
 //        user := c.MustGet(gin.AuthUserKey).(string)

 //        if secret, ok := secrets[user]; ok {
 //            c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
 //        } else {
 //            c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
 //        }
 //    })
}


