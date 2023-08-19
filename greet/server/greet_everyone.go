package main

import (
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)


func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
log.Printf("greet evereyone was invoked")

for {
	req, err :=stream.Recv()
	if err == io.EOF{
		return nil
	}

	if err !=nil{
		log.Fatal("Error while rreading client stream: %v\n", err)
	}

	res :="Hello" +req.FirstName + "!"
	err = stream.Send(&pb.GreetResponse{
		Result: res,
	})

	if err !=nil{
		log.Fatalf("error while sending data to client: %v\n", err)
	}
}
}
