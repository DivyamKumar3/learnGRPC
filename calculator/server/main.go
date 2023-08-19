package main

import (
	"log"
	"net"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
)

var address string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {

	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("failed to listen on : %v\n", err)
	}

	log.Printf("listening on %s\n", address)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s,&Server{})
	
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
