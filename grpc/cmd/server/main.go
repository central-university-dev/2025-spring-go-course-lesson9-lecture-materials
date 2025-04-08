package main

import (
	"net"

	"gitlab.tcsbank.ru/dealer/toolbox/edu9/grpc/internal"
	"gitlab.tcsbank.ru/dealer/toolbox/edu9/grpc/stream"
	"gitlab.tcsbank.ru/dealer/toolbox/edu9/grpc/unary"
	"google.golang.org/grpc"
)

func main() {

	stackSvc := internal.NewStackServer(internal.NewStack())

	server := grpc.NewServer()
	unary.RegisterIntStackServer(server, stackSvc)
	stream.RegisterIntStreamServer(server, internal.NewStreamingService())

	lis, _ := net.Listen("tcp", ":50051")
	_ = server.Serve(lis)
}
