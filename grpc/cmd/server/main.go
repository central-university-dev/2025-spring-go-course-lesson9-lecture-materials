package main

import (
	"net"

	"google.golang.org/grpc"
	"lecture9.demo/grpc/internal"
	"lecture9.demo/grpc/stream"
	"lecture9.demo/grpc/unary"
)

func main() {

	stackSvc := internal.NewStackServer(internal.NewStack())

	server := grpc.NewServer()
	unary.RegisterIntStackServer(server, stackSvc)
	stream.RegisterIntStreamServer(server, internal.NewStreamingService())

	lis, _ := net.Listen("tcp", ":50051")
	_ = server.Serve(lis)
}
