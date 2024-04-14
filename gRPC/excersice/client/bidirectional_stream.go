package main

import (
	"context"
	pb "excersice/excersice/proto"
	"io"
	"log"
)

func InvokeBidrectionalStream(c pb.ExcersiceServiceClient) {
	stream, err := c.BiDirectionalStreamAvg(context.Background())

	numbers := []int32{1, 5, 3, 6, 2, 20}

	if err != nil {
		log.Fatalf("Error while initializing stream")
	}

	block := make(chan struct{})
	log.Printf("Client will send %v in order\n", numbers)
	// Client -> Server Go Routine
	go func() {

		// Send all of the number
		for _, req := range numbers {
			clientErr := stream.Send(&pb.BidirectionalRequest{
				Num: req,
			})

			// If error while sending data to client
			if clientErr != nil {
				log.Printf("Error while send data to stream: %v\n", clientErr)
			}

		}

		// Close client side stream
		stream.CloseSend()
	}()

	// Server -> Client
	go func() {
		for {
			recv, serverErr := stream.Recv()

			// If serverside stream end
			if serverErr == io.EOF {
				break
			}

			// Error while receiving message
			if serverErr != nil {
				log.Fatalf("Error while receiving message from server: %v", serverErr)
			}

			log.Printf("Received: %v\n", recv.MaximumNumber)
		}

		// Shutdown channel
		close(block)
	}()

	// Set block state while channel close
	<-block
}
