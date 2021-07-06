package movie

import (
	"encoding/json"
	movie "ticket-manager/rpc/grpc/protos/movie"
	"ticket-manager/service"

	"golang.org/x/net/context"
)

type MoiveRPCServer struct{}

func (s *MoiveRPCServer) GetAllMovies(ctx context.Context, request *movie.MovieMsgRequest) (*movie.MovieMsgReply, error) {
	var movieSerive service.MovieService
	movies, err := movieSerive.FindAllMovies()
	if err != nil {
		return &movie.MovieMsgReply{Message: "err: " + err.Error()}, nil
	}
	josnStr, _ := json.Marshal(movies)
	return &movie.MovieMsgReply{Message: string(josnStr)}, nil
}
