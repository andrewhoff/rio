package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/andrewhoff/rio/helloproto"
	"google.golang.org/grpc"
)

const port = ":8111"

func Run() {
	fmt.Print("\x1b[33;1mStarting Server...\n")

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &helloServer{})

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	panic(s.Serve(lis))
}

type helloServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *helloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("\x1b[33;1mSaying Hello to the %s...\n", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
