package main

import (
	"context"
	pb "github.com/ronaldoafonso/rbox/gcommand"
	"github.com/ronaldoafonso/rbox/rbox"
	"strings"
)

type server struct {
	pb.UnimplementedRemoteCommandServer
}

func (s *server) Get(c context.Context, r *pb.GetRequest) (*pb.ReturnMsg, error) {
	b := rbox.NewRBox(r.GetBoxname())
	returnMsg := "OK: "

	switch r.GetField() {
	case "ssid":
		SSIDs, err := b.GetSSIDs()
		if err != nil {
			return &pb.ReturnMsg{ReturnMsg: err.Error()}, err
		}
		returnMsg += strings.Join(SSIDs, " ")
	case "macs":
		MACs, err := b.GetMACs()
		if err != nil {
			return &pb.ReturnMsg{ReturnMsg: err.Error()}, err
		}
		returnMsg += strings.Join(MACs, " ")
	default:
		returnMsg += "Field not implemented."
	}

	return &pb.ReturnMsg{ReturnMsg: returnMsg}, nil
}

func (s *server) Set(c context.Context, r *pb.SetRequest) (*pb.ReturnMsg, error) {
	b := rbox.NewRBox(r.GetBoxname())
	returnMsg := "OK: "

	switch r.GetField() {
	case "ssid":
		err := b.SetSSIDs(r.GetValue())
		if err != nil {
			return &pb.ReturnMsg{ReturnMsg: err.Error()}, err
		}
	case "macs":
		err := b.SetMACs(r.GetValue())
		if err != nil {
			return &pb.ReturnMsg{ReturnMsg: err.Error()}, err
		}
	default:
		returnMsg = "Field not implemented."
	}

	return &pb.ReturnMsg{ReturnMsg: returnMsg}, nil
}
