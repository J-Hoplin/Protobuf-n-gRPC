package main

import (
	"context"
	pb "excersice/excersice/proto"
	"log"
)

func (s *Server) UnarySum(ctx context.Context, req *pb.UnarySumReqest) (*pb.UnarySumResponse, error) {
	log.Println("Server side 'UnarySum' invoked ")
	return &pb.UnarySumResponse{Result: req.Num1 + req.Num2}, nil
}
