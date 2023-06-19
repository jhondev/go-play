package server

import (
	"context"
	"fmt"
	"grpc-patterns/proto"

	"google.golang.org/grpc"
)

type GrpcService struct{}

func New() *grpc.Server {
	svr := grpc.NewServer()
	proto.RegisterServiceServer(svr, &GrpcService{})
	return svr
}

// UpdateOptional implements proto.ServiceServer.
func (*GrpcService) UpdateOptional(ctx context.Context, req *proto.OptionalRequest) (*proto.OptionalResponse, error) {
	fmt.Println("REQUEST VALUES")
	fmt.Print(req)
	return &proto.OptionalResponse{}, nil
}
