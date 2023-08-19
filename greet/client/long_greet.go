package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient){
	log.Println("doLongGreet was invoked")

	reqs :=[]*pb.GreetRequest{
		{FirstName: "Divyam"},
		{FirstName: "kumar"},
		{FirstName: "gupta"},
	}
	stream, err :=c.LongGreet(context.Background())

	if err !=nil{
		log.Fatal("error while calling longgreet %v\n", err)
	}

	for _, req:=range reqs {
		log.Printf("sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err !=nil{
		log.Fatalf("Error while receiving response from longGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}