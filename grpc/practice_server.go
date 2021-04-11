package main

import (
	"fmt"
	proto "go-business/grpc/proto/practice"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

type practiceServer struct {
	proto.UnimplementedPracticeServer
	mu sync.Mutex
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (s *practiceServer) MsgNotice(req *proto.MsgNoticeRequest, stream proto.Practice_MsgNoticeServer) error {
	// 随机获取遍历次数
	num := rand.Intn(20)
	for i := 0; i < num; i++ {
		msg := fmt.Sprintf("The msg num is %d", num)
		rsp := &proto.MsgNoticeResponse{Title: "server push", Desc: msg}
		err := stream.Send(rsp)
		if err != nil {
			fmt.Println("send err: ", err)
			return err
		}
	}
	return nil
}

// TODO: 断点续传
func (s *practiceServer) UploadFile(stream proto.Practice_UploadFileServer) error {
	fileName := "practice-" + time.Now().Format("2006-01-02")
	fileSize := int32(0)
	startTime := time.Now()
	for {
		req, err := stream.Recv()
		// 读完
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&proto.UploadFileResponse{
				FileName:    fileName,
				FileSize:    fileSize,
				ElapsedTime: int32(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			fmt.Println("recv data err: ", err)
			return err
		}
		fileSize += 1
		fmt.Printf("recv data, name: %s, phone: %s, work: %s ", req.Name, req.Phone, req.Work)
	}
}

func (s *practiceServer) SimpleChat(stream proto.Practice_SimpleChatServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		rMsg := "Hello " + msg.Msg
		err = stream.Send(&proto.ChatMsg{Msg: rMsg})
		if err != nil {
			return err
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":52002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterPracticeServer(s, &practiceServer{})
	fmt.Println("grpc server listen in 52002...")
	s.Serve(listener)
}
