package main

import (
	"context"
	pb "excersice/excersice/proto"
	"fmt"
	"log"
)

func InvokeClientStream(c pb.ExcersiceServiceClient) {

	reqs := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	stream, err := c.ClientStreamAvg(context.Background())

	if err != nil {
		log.Fatalf("Error while getting client stream: %v", err)
	}

	for _, num := range reqs {
		stream.Send(&pb.ClientStreamRequest{
			Num: num,
		})
	}
	// Make client stream close
	res, err := stream.CloseAndRecv()
	fmt.Println("Double of Average: ", res)
}
