package services

import (
	"context"

	"github.com/sandeep-cs-dev/grpc/server/proto"
)

type MathServiceServer struct {
	proto.UnimplementedMathServiceServer
}

func (mathservice *MathServiceServer) Add(
	ctx context.Context, addrequest *proto.AddRequest) (*proto.AddResponse, error) {
	return &proto.AddResponse{
		Sum: (addrequest.Number1 + addrequest.Number2),
	}, nil
}
