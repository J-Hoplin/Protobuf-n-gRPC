package main

import (
	pb "excersice/excersice/proto"
)

func (s *Server) ServerStreamPrimeNumber(input *pb.ServerStreamRequest, stream pb.ExcersiceService_ServerStreamPrimeNumberServer) error {
	// Bind client side request input
	n := input.Num

	// Prime number start from 2
	var k int32 = 2
	// Calculate Prime Number decomposition
	for {
		// Loop end condition
		if n <= 1 {
			break
		}

		// If it's remainder is 0,
		if n%k == 0 {
			stream.Send(&pb.ServerStreamResponse{
				PrimeDecomposition: k,
			})
			n = n / k
		} else {
			k++
		}
	}

	// End stream of the server
	return nil
}
