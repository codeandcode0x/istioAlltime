package user

import (
	"encoding/json"
	"golang.org/x/net/context"
	user "ticket-manager/rpc/grpc/protos/user"
	"ticket-manager/service"
)

type UserRpcServer struct {}

func (s *UserRpcServer) GetAllUsers(ctx context.Context, request *user.UserMsgRequest) (*user.UserMsgReply, error) {
	var userSerive  service.UserService
	users, err := userSerive.FindAllUsers()
	if err != nil {
		return &user.UserMsgReply{Message: "err: " + err.Error()}, nil
	}
	josnStr,_ := json.Marshal(users)
	return &user.UserMsgReply{Message: "your user is: " + string(josnStr)}, nil
}