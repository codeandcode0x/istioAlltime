//main
package main

import (
	"fmt"
	"net"
	"ticket-manager/route"
	movie "ticket-manager/rpc/grpc/protos/movie"
	user "ticket-manager/rpc/grpc/protos/user"
	movieRpc "ticket-manager/rpc/grpc/service/movie"
	userRpc "ticket-manager/rpc/grpc/service/user"

	tracing "github.com/codeandcode0x/traceandtrace-go"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	endless.ListenAndServe(":8080", r)

	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        http.TimeoutHandler(r, time.Second*5, ""),
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()

}

func RpcServer() {
	// add gRPC tracing
	rpcOption, closer, _ := tracing.AddRpcServerTracing("RpcServer")
	defer closer.Close()

	lis, err := net.Listen("tcp", ":22530")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer(rpcOption)
	movie.RegisterMovieRPCServer(s, &movieRpc.MoiveRPCServer{})
	user.RegisterUserRPCServer(s, &userRpc.UserRPCServer{})
	reflection.Register(s)
	err = s.Serve(lis)

	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
