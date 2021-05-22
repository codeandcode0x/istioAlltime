//main
package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"ticket-manager/route"
)

func main() {

	r := gin.Default()
	//s := &http.Server{
	//	Addr:           ":8080",
	//	Handler:        r,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	route.DefinitionRoute(r)
	endless.ListenAndServe(":8080", r)
}






