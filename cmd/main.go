package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "test1/gen"
	"test1/internal/test_service/handlers"
)

func main() {
	var opts []grpc.ServerOption
	handler := handlers.NewTestHandler()
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(opts...)
	pb.RegisterTestServiceServer(server, handler)

	fmt.Println("server run on", "localhost:50051")
	server.Serve(lis)
}
