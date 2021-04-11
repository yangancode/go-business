package main

import (
	"context"
	"fmt"
	"go-business/grpc/proto/greeter"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, req *greeter.HelloRequest) (rsp *greeter.HelloReply, err error) {
	rsp = &greeter.HelloReply{Message: "Hello " + req.Name}
	return rsp, nil
}

func main() {
	listener, err := net.Listen("tcp", ":52001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, &server{})

	//reflection.Register(s)
	fmt.Println("gRPC server listen in 52001...")
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
