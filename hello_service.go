package main

import (
	"context"
	pb "github.com/NFCHKK/grpc/proto"
	"log"
)

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (out *pb.HelloReply, err error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SaveNotes(ctx context.Context, in *pb.SaveNotesRequest) (out *pb.SaveNotesResponse, err error) {
	log.Printf("Received: %+v", *in)
	return &pb.SaveNotesResponse{
		Head: &pb.RspHead{
			Code: 0, Msg: "success",
		},
	}, nil
}

func (s *server) GetNotes(ctx context.Context, in *pb.GetNotesRequest) (out *pb.GetNotesResponse, err error) {
	log.Printf("Received: %+v", *in)
	return &pb.GetNotesResponse{
		Head: &pb.RspHead{
			Code: 0, Msg: "success",
		},
		Notes: "this is a serial notes",
	}, nil
}
