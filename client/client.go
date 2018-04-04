package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/andrewhoff/rio/helloproto"
	"google.golang.org/grpc"
)

func Run() {
	fmt.Print("\x1b[32;1mStarting Client...\n")

	// Set up a connection to the server.
	conn, err := grpc.Dial(":8111", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Boldies"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Greeting: %s\n", r.Message)
}
