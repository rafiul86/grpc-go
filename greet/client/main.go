package main

import (
	"log"

	pb "github.com/rafiul86/grpc-go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
	defer conn.Close()
}