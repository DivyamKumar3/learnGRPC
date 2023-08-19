package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)
var address string ="localhost:50551"

func main(){
	conn, err :=grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err !=nil{
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c:= pb.NewGreetServiceClient(conn)

	//doGreet(c)
	//doGrretManyTimes(c)
	//doLongGreet(c)
	doGreetEveryone(c)
}