package handlers

import (
	"context"
	pb "test1/gen"
)

type TestHandler struct {
	pb.UnimplementedTestServiceServer
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (handler *TestHandler) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello, " + req.Name,
	}, nil
}
