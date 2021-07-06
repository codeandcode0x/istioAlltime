package user

import (
	"context"
	"fmt"
	user "ticket-manager/rpc/grpc/protos/user"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":22530", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := user.NewUserRPCClient(conn)
	r, err := c.GetAllUsers(context.Background(), &user.UserMsgRequest{Count: 100})
	if err != nil {
		fmt.Printf("could not request: %v", err)
	}

	fmt.Printf("get user count : %s !\n", r.Message)
}
