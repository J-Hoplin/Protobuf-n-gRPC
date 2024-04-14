package main

import (
	"context"
	pb "excersice/excersice/proto"
	"fmt"
	"io"
	"log"
)

func InvokeServerStream(c pb.ExcersiceServiceClient) {

	req := &pb.ServerStreamRequest{
		Num: 210,
	}
	stream, err := c.ServerStreamPrimeNumber(context.Background(), req)
	if err != nil {
		log.Fatalf("Fail to create stream: %v", err)
	}
	fmt.Println("Stream created")

	// Initialize with empty array to prevent unused buffer
	primeNumbers := make([]int32, 0)
	for {
		recv, err := stream.Recv()

		// If server stream ended
		if err == io.EOF {
			break
		}

		// If it's error is not an io.EOF -> Fail while streaming
		if err != nil {
			log.Fatalf("Error while receving message: %v", err)
			return
		}

		primeNumbers = append(primeNumbers, recv.PrimeDecomposition)
	}

	for index, element := range primeNumbers {
		log.Printf("%v. %v", index, element)
	}
}
