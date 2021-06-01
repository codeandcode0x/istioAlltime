package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	movie "ticket-manager/rpc/grpc/protos/movie"
	user "ticket-manager/rpc/grpc/protos/user"
)

func main()  {
	conn, err := grpc.Dial(":20153", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := movie.NewMovieRPCClient(conn)
	r, err := c.GetAllMovies(context.Background(), &movie.MovieMsgRequest{Count: 100})
	if err != nil {
		fmt.Printf("could not request: %v", err)
	}

	c1 := user.NewUserRPCClient(conn)
	r1, err := c1.GetAllUsers(context.Background(), &user.UserMsgRequest{Count: 100})

	fmt.Printf("movie list : %s !\n", r.Message)
	fmt.Printf("user list : %s !\n",  r1.Message)
}