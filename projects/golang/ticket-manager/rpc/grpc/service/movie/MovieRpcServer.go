package movie

import (
	"encoding/json"
	"golang.org/x/net/context"
	movie "ticket-manager/rpc/grpc/protos/movie"
	"ticket-manager/service"

	//"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
	//"net"
	//user "ticket-manager/rpc/grpc/protos/user"
	//userRpc "ticket-manager/rpc/grpc/service/user"
)

type MoiveRpcServer struct {}

func (s *MoiveRpcServer) GetAllMovies(ctx context.Context, request *movie.MovieMsgRequest) (*movie.MovieMsgReply, error) {
	var movieSerive  service.MovieService
	movies, err := movieSerive.FindAllMovies(3)
	if err != nil {
		return &movie.MovieMsgReply{Message: "err: " + err.Error()}, nil
	}
	josnStr,_ := json.Marshal(movies)
	return &movie.MovieMsgReply{Message: "your movie is: " + string(josnStr)}, nil
}

//func main()  {
//	lis, err := net.Listen("tcp", ":20153")
//	if err != nil {
//		fmt.Printf("failed to listen: %v", err)
//		return
//	}
//
//	s := grpc.NewServer()
//	movie.RegisterMovieRPCServer(s, &MoiveRpcServer{})
//	user.RegisterUserRPCServer(s, &userRpc.UserRpcServer{})
//	reflection.Register(s)
//	err = s.Serve(lis)
//
//	if err != nil {
//		fmt.Printf("failed to serve: %v", err)
//		return
//	}
//}