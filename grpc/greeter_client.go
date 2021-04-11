package main

import (
	"context"
	"fmt"
	"go-business/grpc/proto/greeter"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	// 发起连接，WithInsecure表示使用不安全的连接，即不使用SSL
	conn, err := grpc.Dial("127.0.0.1:52001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}

	defer conn.Close()

	c := greeter.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	r, err := c.SayHello(ctx, &greeter.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("call service failed: %v", err)
	}
	fmt.Println("call service success: ", r.Message)
}
