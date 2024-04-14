package main

import (
	pb "excersice/excersice/proto"
	"io"
	"log"
)

func (s *Server) BiDirectionalStreamAvg(stream pb.ExcersiceService_BiDirectionalStreamAvgServer) error {
	var currentMaximumNumber int32 = 0
	for {
		recv, err := stream.Recv()

		// If client request is done
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while processing stream: %v", err)
		}
		// If received number value is lower than current Maximum Number, ignore
		if recv.Num < currentMaximumNumber {
			continue
		}

		currentMaximumNumber = recv.Num
		err = stream.Send(&pb.BidirectionalResponse{
			MaximumNumber: currentMaximumNumber,
		})
		if err != nil {
			log.Fatalf("Error while sending data: %v", err)
		}
	}
}
