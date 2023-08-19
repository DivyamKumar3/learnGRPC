package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient){
	log.Printf("do sum was invoked")
	res, err:= c.Sum(context.Background(), &pb.SumRequest{
FirstNumber:1,
SecondNumber:2,
	})

if err!=nil{
	log.Fatal("could not sum: %v", err)
}
log.Printf("sum: %d\n", res.Result)
}