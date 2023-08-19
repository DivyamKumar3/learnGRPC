package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doGrretManyTimes(c pb.GreetServiceClient){
	log.Println("doGretMany times was invoked")

	req :=&pb.GreetRequest{
		FirstName:"Divyam",
	}
stream , err :=c.GreetManyTimes(context.Background(), req)

if err !=nil{
	log.Fatal("error while calling greet many times: %v\n", err)

}

for {
	msg, err:=stream.Recv()
	if err==io.EOF{
		break
	}
	if err !=nil{
		log.Fatal("error while reading the stream: %v/n", err)
	}
	log.Printf("GreetManyTimes: %s\n", msg.Result)
}





}
