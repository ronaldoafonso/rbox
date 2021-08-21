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

	// Get
	for _, field := range []string{"ssid", "macs"} {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		data := &pb.GetRequest{
			Boxname: "788a20298f81.z3n.com.br",
			Field:   field,
		}
		if retMessage, err := c.Get(ctx, data); err != nil {
			log.Fatalf("could not get %v: %v.", field, err)
		} else {
			log.Printf("Get %v: %v.", field, retMessage.GetReturnMsg())
		}
	}

	// Set
	data := []struct {
		field string
		value string
	}{
		{"ssid", "afonso ronaldo afonso"},
		{"macs", "11:11:11:11:11:11 22:22:22:22:22:22 33:33:33:33:33:33"},
	}
	for _, d := range data {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		sdata := &pb.SetRequest{
			Boxname: "788a20298f81.z3n.com.br",
			Field:   d.field,
			Value:   d.value,
		}
		if retMessage, err := c.Set(ctx, sdata); err != nil {
			log.Fatalf("could not set %v->%v: %v.", d.field, d.value, err)
		} else {
			log.Printf("Set %v->%v: %v.", d.field, d.value, retMessage.GetReturnMsg())
		}
	}
}
