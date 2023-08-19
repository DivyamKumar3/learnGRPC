package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Printf("do avaegarge was invoked")


	stream, err :=c.Avg(context.Background())


	if err !=nil{
		log.Fatal("error while opening the stream : %v\n", err)
	}

	numbers :=[]int32{3,5,9,54,23}

	for _, number:=range numbers {
		log.Printf("sending number: %d\n", number)
		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}

	res, err :=stream.CloseAndRecv()

	if err!=nil{
		log.Fatal("error while receiving response: %v\n", err)
	}

	log.Printf("avg: %f\n", res.Result)
}