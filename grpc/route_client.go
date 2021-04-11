package main

import (
	"context"
	"flag"
	proto "go-business/grpc/proto/route"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func printFeatures(client proto.RouteGuideClient, rect *proto.Rectangle) {
	log.Printf("Looking for features within %v", rect)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.ListFeatures(ctx, rect)
	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Printf("Feature: name: %q, point:(%v, %v)", feature.GetName(),
			feature.GetLocation().GetLatitude(), feature.GetLocation().GetLongitude())
	}
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := proto.NewRouteGuideClient(conn)
	printFeatures(client, &proto.Rectangle{
		Lo: &proto.Point{Latitude: 400000000, Longitude: -750000000},
		Hi: &proto.Point{Latitude: 420000000, Longitude: -730000000},
	})
}
