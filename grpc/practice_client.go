package main

import (
	"context"
	"fmt"
	proto "go-business/grpc/proto/practice"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
	"time"
)

func recvMsg(client proto.PracticeClient) {
	log.Printf("start recv")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &proto.MsgNoticeRequest{
		Uid:        10000,
		LatestTime: "2021-04-01 00:00:00",
	}
	stream, err := client.MsgNotice(ctx, req)
	if err != nil {
		log.Fatal("err: ", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("err: ", err)
		}
		log.Printf("server msg: %s", msg.Desc+msg.Title)
	}
}

func uploadFile(client proto.PracticeClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.UploadFile(ctx)
	if err != nil {
		log.Fatal("err: ", err)
	}
	for i := 0; i < 10; i++ {
		data := &proto.UploadFileRequest{Name: "yangan", Phone: "123456", Work: "none"}
		err := stream.Send(data)
		if err != nil {
			log.Fatal("err: ", err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("err: ", err)
	}
	fmt.Printf("rsp, name: %s, size: %d, time: %d", reply.FileName, reply.FileSize, reply.ElapsedTime)
}

func simpleChat(client proto.PracticeClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.SimpleChat(ctx)
	if err != nil {
		log.Fatal("err: ", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a msg : %v", err)
			}
			log.Printf("Got message %s", in.Msg)
		}
	}()

	for i := 0; i < 10; i++ {
		msg := &proto.ChatMsg{Id: uint32(i), Msg: "chat: " + strconv.Itoa(i)}
		err := stream.Send(msg)
		if err != nil {
			log.Fatalf("Failed to send a msg : %v", err)
		}
	}

	stream.CloseSend()
	<-waitc
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:52002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := proto.NewPracticeClient(conn)
	//recvMsg(client)

	//uploadFile(client)

	simpleChat(client)

}
