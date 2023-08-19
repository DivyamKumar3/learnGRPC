package main

import (
	"context"
	"log"
	"io"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto";
)

func doPrimes(c pb.CalculatorServiceClient){
log.Println("doprime was invoked")


req :=&pb.PrimeRequest{
	Number:12390392840,
}

stream, err :=c.Primes(context.Background(), req)

if err !=nil{
	log.Fatal("error while calling primes : %v\n", err)
}

for{
	res, err:=stream.Recv()

	if err ==io.EOF{
		break
	}

	if err!=nil{
		log.Fatal("error while reading stream: %v\n", err)
	}


	log.Printf("primes:%d\n", res.Result)
}

}