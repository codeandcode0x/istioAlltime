package route

import (
	"github.com/gin-gonic/gin"
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
	router.LoadHTMLGlob("web/*")
	// user
	var userController *controller.UserController
	router.GET("/api/users", userController.GetAllUsers)
	router.POST("/api/user/create", userController.CreateUser)
	router.POST("/api/user/update", userController.UpdateUser)
	router.POST("/api/user/delete", userController.DeleteUser)

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


