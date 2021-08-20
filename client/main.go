package main

import (
	"context"
	pb "github.com/ronaldoafonso/rbox/gcommand"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("fail to connect: %v.", err)
	}
	defer conn.Close()

	c := pb.NewRemoteCommandClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Get
	getSSID := &pb.GetRequest{
		Boxname: "boxname",
		Field:   "ssid",
	}
	if retMessage, err := c.Get(ctx, getSSID); err != nil {
		log.Fatalf("could not get SSID: %v.", err)
	} else {
		log.Printf("Get SSID: %v.", retMessage.GetReturnMsg())
	}

	// Set
	setSSID := &pb.SetRequest{
		Boxname: "boxname",
		Field:   "ssid",
		Value:   "z3n",
	}
	if retMessage, err := c.Set(ctx, setSSID); err != nil {
		log.Fatalf("could not set SSID: %v.", err)
	} else {
		log.Printf("Set SSID: %v.", retMessage.GetReturnMsg())
	}
}
