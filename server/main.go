package main

import (
	"github.com/ronaldoafonso/rbox/gcommand"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error listening: %v.", err)
	}

	s := grpc.NewServer()
	gcommand.RegisterRemoteCommandServer(s, &server{})

	log.Printf("Starting GRPC server!")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v.", err)
	}
}
