package main

import (
	pb "excersice/excersice/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var host string = ":8080"

// Deletgate top process of ServiceServer
type Server struct {
	pb.ExcersiceServiceServer
}

func main() {
	listener, err := net.Listen("tcp", host)

	if err != nil {
		log.Fatalf("Fail to generate listenr: %v\n", err)
	}

	log.Printf("Listening on %v\n", host)

	// New server
	server := grpc.NewServer()
	// Server register interface
	serverRegister := new(Server)
	pb.RegisterExcersiceServiceServer(server, serverRegister)

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Fail to serve: %v\n", err)
	}
}
