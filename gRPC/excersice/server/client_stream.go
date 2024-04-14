package main

import (
	pb "excersice/excersice/proto"
	"io"
	"log"
)

// Get average of numbers from client and double it
func (s *Server) ClientStreamAvg(stream pb.ExcersiceService_ClientStreamAvgServer) error {
	numbers := make([]int32, 0)
	var sum int32 = 0
	for {
		// 내부적으로 메세지가 오기 전까지 blocking state로 유지한다.
		recv, err := stream.Recv()

		// If error is EOF -> End stream from client
		if err == io.EOF {
			break
		}

		// If it's error without EOF
		if err != nil {
			log.Fatalf("Error in processing stream: %v", err)
		}
		numbers = append(numbers, recv.Num)
	}

	for _, num := range numbers {
		sum += num
	}

	avg := sum / int32(len(numbers))
	log.Printf("Integer Average value from client is %d\n", avg)

	// Send message to client and close connection
	stream.SendAndClose(&pb.ClientStreamResponse{
		Avg: avg * 2,
	})
	return nil
}
