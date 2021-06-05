package movie

import (
	"encoding/json"
	movie "ticket-manager/rpc/grpc/protos/movie"
	"ticket-manager/service"

	"golang.org/x/net/context"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
	//"net"
	//user "ticket-manager/rpc/grpc/protos/user"
	//userRpc "ticket-manager/rpc/grpc/service/user"
)

type MoiveRpcServer struct{}

func (s *MoiveRpcServer) GetAllMovies(ctx context.Context, request *movie.MovieMsgRequest) (*movie.MovieMsgReply, error) {
	var movieSerive service.MovieService
	movies, err := movieSerive.FindAllMovies(3)
	if err != nil {
		return &movie.MovieMsgReply{Message: "err: " + err.Error()}, nil
	}
	josnStr, _ := json.Marshal(movies)
	return &movie.MovieMsgReply{Message: string(josnStr)}, nil
}
