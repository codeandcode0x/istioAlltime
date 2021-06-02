//main
package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"ticket-manager/route"
	movie "ticket-manager/rpc/grpc/protos/movie"
	user "ticket-manager/rpc/grpc/protos/user"
	movieRpc "ticket-manager/rpc/grpc/service/movie"
	userRpc "ticket-manager/rpc/grpc/service/user"
)

var (
	g errgroup.Group
)

func main() {
	r := gin.Default()
	route.DefinitionRoute(r)
	// rpc goroutine
	g.Go(func() error {
		RpcServer()
		return nil
	})

	endless.ListenAndServe(":8081", r)
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






