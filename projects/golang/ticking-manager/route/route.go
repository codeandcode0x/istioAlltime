package route

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ticking-manager/src"
	"time"
)

func DefinitionRoute(router *gin.Engine) {

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 模拟一些私人数据
	var secrets = gin.H{
	    "admin":    gin.H{"email": "admin@bar.com", "phone": "123433"},
	    "demo": gin.H{"email": "demo@example.com", "phone": "666"},
	}

	router.GET("/", src.Home)
	router.GET("/api/version", src.GetVersion)

	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin", 
		"demo":  "demo", 
	}))

	authorized.Use(AuthRequired())
	{
		 authorized.GET("/user", func(c *gin.Context) {
	        // 获取用户，它是由 BasicAuth 中间件设置的
	         c.JSON(http.StatusOK, gin.H{"OK": "OK"})
	    })

	    authorized.GET("/logout", func(c *gin.Context) {
	        // 获取用户，它是由 BasicAuth 中间件设置的
	         c.JSON(http.StatusOK, gin.H{"OK": "OK"})
	    })
	}

	authorized.GET("/secrets", func(c *gin.Context) {
        // 获取用户，它是由 BasicAuth 中间件设置的
        user := c.MustGet(gin.AuthUserKey).(string)

        if secret, ok := secrets[user]; ok {
            c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
        } else {
            c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
        }
    })
}


func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()
    
        // 设置 example 变量
        c.Set("example", "12345")
    
        // 请求之前...
    
        c.Next()
    
        // 请求之后...
        latency := time.Since(t)
        // 打印请求处理时间
        log.Print(latency)
    
        // 访问即将发送的响应状态码
        status := c.Writer.Status()
        log.Println(status)
    }
}



