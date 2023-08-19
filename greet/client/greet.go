package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient){
	log.Println("doGreet was invoked")
	res, err :=c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Divyam",
	})

	if err !=nil{
log.Fatalf("could not greet: %v\n", err)
	}

	log.Printf("greeting : %s\n", res.Result)
}