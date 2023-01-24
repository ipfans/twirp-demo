package greetersvc

import (
	"context"

	pb "github.com/ipfans/twirp-demo/rpc/greeter/v1"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{
		Message: "Hello, " + req.Name,
	}, nil
}
