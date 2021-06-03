package middleware

import (
	//"github.com/gin-contrib/timeout"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func TimeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		//ch := make(chan int)
		//ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		//go func(ctx context.Context) {
		//	if ctx.Err() == context.DeadlineExceeded {
		//		log.Println("request timeout :"+ ctx.Err().Error())
		//		c.Writer.WriteHeader(http.StatusGatewayTimeout)
		//		ch <- 1
		//	}
		//}(ctx)
		//
		//select {
		//case status:= <-ch:
		//	log.Println("request timeout2 .....", status)
		//	c.JSON(http.StatusGatewayTimeout,  gin.H{
		//		"code": http.StatusGatewayTimeout,
		//		"error": "request timeout",
		//	})
		//	c.Request = c.Request.WithContext(ctx)
		//	c.Next()
		//	c.Abort()
		//case <-ctx.Done():
		//	cancel()
		//
		//	return
		//}
		ch := make(chan int, 1)
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func(c *gin.Context) {
			if ctx.Err() == context.DeadlineExceeded {
				log.Println("request timeout1 .... :" + ctx.Err().Error())
				ch <- 1
			}
			cancel()
			log.Println("request timeout2 .... :")
		}(c)

		for {
			select {
			case result := <-ch:
				log.Println("request timeout3 .... :", result)
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.JSON(http.StatusGatewayTimeout, gin.H{
					"code":  http.StatusGatewayTimeout,
					"error": "request timeout",
				})
				c.Abort()
				return
				// case <-ctx.Done():
				// 	c.Request = c.Request.WithContext(ctx)
				// 	c.Next()
				// 	log.Println("request timeout4 .... :")
				// 	return
			}
		}
	}
}

func TimedHandler(duration time.Duration) func(c *gin.Context) {
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
				body: gin.H{
					"code":  http.StatusGatewayTimeout,
					"error": "request timeout",
				},
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
			c.Writer.WriteHeader(http.StatusGatewayTimeout)
			c.JSON(res.status, res.body)
			c.Abort()
			return
		}
	}
}
