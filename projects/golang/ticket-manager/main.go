//main
package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"ticket-manager/route"
	movie "ticket-manager/rpc/grpc/protos/movie"
	user "ticket-manager/rpc/grpc/protos/user"
	movieRpc "ticket-manager/rpc/grpc/service/movie"
	userRpc "ticket-manager/rpc/grpc/service/user"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
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

	g.Go(func() error {
		RpcServer()
		return nil
	})

	endless.ListenAndServe(":8080", r)
}


func RpcServer() {
	lis, err := net.Listen("tcp", ":20153")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	movie.RegisterMovieRPCServer(s, &movieRpc.MoiveRpcServer{})
	user.RegisterUserRPCServer(s, &userRpc.UserRpcServer{})
	reflection.Register(s)
	err = s.Serve(lis)

	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}






