package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient){
	log.Printf("doGreetEveryone was invoked")

	stream, err :=c.GreetEveryone(context.Background())

	if err!=nil{
		log.Fatal("Error while creating stream: %v\n", err)
	}

	reqs :=[]*pb.GreetRequest{
		{FirstName: "Divyam"},
		{FirstName: "kumar"},
		{FirstName: "gupta"},
	}

	waitc :=make(chan struct{})

	go func(){
		for _, req :=range reqs {
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func(){
    for{
		res, err :=stream.Recv()

		if err ==io.EOF{
			break
		}

		if err !=nil{
			log.Printf("Error while receiving: %v\n", err)
			break
		}
		log.Printf("received : %v\n", res.Result)
	}
	close(waitc)
	}()

	<-waitc
}