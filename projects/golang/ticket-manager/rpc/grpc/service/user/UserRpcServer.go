package user

import (
	"encoding/json"
	user "ticket-manager/rpc/grpc/protos/user"
	"ticket-manager/service"

	"golang.org/x/net/context"
)

type UserRPCServer struct{}

func (s *UserRPCServer) GetAllUsers(ctx context.Context, request *user.UserMsgRequest) (*user.UserMsgReply, error) {
	var userSerive service.UserService
	users, err := userSerive.FindAllUsers()
	if err != nil {
		return &user.UserMsgReply{Message: "err: " + err.Error()}, nil
	}
	josnStr, _ := json.Marshal(users)
	return &user.UserMsgReply{Message: "your user is: " + string(josnStr)}, nil
}
