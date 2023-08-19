package main

import (
    "log"
	"fmt"
	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)
func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error{
	log.Printf("greeting many times function was invoked : %v\n", in)

  for i :=0; i<10; i++{
	res :=fmt.Sprintf("Helloo %s, number %d", in.FirstName, i)
	stream.Send(&pb.GreetResponse{
		Result: res,
	})
  }
  return nil
}
