package main

import (
	"context"
	pb "excersice/excersice/proto"
	"log"
)

func InvokeUnarySum(client pb.ExcersiceServiceClient) {
	res, err := client.UnarySum(context.Background(), &pb.UnarySumReqest{
		Num1: 10,
		Num2: 20,
	})

	if err != nil {
		log.Fatalf("Fail to invoke RPC endpoint: %v\n", err)
	}

	log.Printf("Invoked: %v\n", res.Result)
}
