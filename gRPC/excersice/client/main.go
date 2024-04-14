package main

import (
	pb "excersice/excersice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var host string = ":8080"

func main() {
	// Ignore http2 default SSL encryption
	// DO NOT APPLY THIS IN PRODUCTION
	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Fail to connect: %v\n", err)
	}

	// Make sure connection to be close after invoke
	defer conn.Close()
	c := pb.NewExcersiceServiceClient(conn)

	// Invoke Unary Sum

	//InvokeUnarySum(c)
	//InvokeClientStream(c)
	//InvokeServerStream(c)
	InvokeBidrectionalStream(c)
}
