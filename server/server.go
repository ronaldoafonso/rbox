package main

import (
	"context"
	pb "github.com/ronaldoafonso/rbox/gcommand"
	"log"
)

type server struct {
	pb.UnimplementedRemoteCommandServer
}

func (s *server) Get(c context.Context, r *pb.GetRequest) (*pb.ReturnMsg, error) {
	log.Printf("Executing GetRequest(%v).", r)
	return &pb.ReturnMsg{ReturnMsg: "Ok Get gRPC"}, nil
}

func (s *server) Set(c context.Context, r *pb.SetRequest) (*pb.ReturnMsg, error) {
	log.Printf("Executing SetRequest(%v).", r)
	return &pb.ReturnMsg{ReturnMsg: "Ok Set gRPC"}, nil
}
