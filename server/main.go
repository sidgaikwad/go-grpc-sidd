package main

import (
	"log"
	"net"

	pb "github.com/sidgaikwad/go-grpc-sidd/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type HelloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &HelloServer{})
	log.Printf("server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
