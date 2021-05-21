//main
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/fvbock/endless"
	"ticking-manager/models"
	"ticking-manager/route"
	"ticking-manager/src/middleware"
	// "net/http"
	// "time"
)

func init() {
	models.Setup()
}

func main() {

	r := gin.Default()
	r.Use(middleware.Logger())
	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        r,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	route.DefinitionRoute(r)
	endless.ListenAndServe(":8080", r)
}






