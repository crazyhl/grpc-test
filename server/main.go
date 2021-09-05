package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"test/pb"
)

type HelloServer struct {
	pb.UnimplementedGreeterServer
}

func (s HelloServer) SayHello(context.Context, *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "test",
	}, nil
}

type OtherServer struct {
	pb.UnimplementedOtherServer
}

func (s OtherServer) SayOther(_ context.Context, in *pb.OtherRequest) (*pb.OtherReply, error) {
	return &pb.OtherReply{
		Message: "Other " + in.Name,
	}, nil
}


var port int = 50051

func newServer() *HelloServer {
	s := &HelloServer{}
	return s
}

func newOtherServer() *OtherServer {
	s := &OtherServer{}
	return s
}

func main()  {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(grpcServer, newServer())
	pb.RegisterOtherServer(grpcServer, newOtherServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		return
	}
}