package main

import (
	"context"
	"github.com/NFCHKK/grpc/db"
	pb "github.com/NFCHKK/grpc/proto"
	"log"
)

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (out *pb.HelloReply, err error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SaveNotes(ctx context.Context, in *pb.SaveNotesRequest) (out *pb.SaveNotesResponse, err error) {
	log.Printf("Received: %+v", *in)
	if in.GetHead() == nil {
		return &pb.SaveNotesResponse{
			Head: &pb.RspHead{
				Code: -1, Msg: "req head empty",
			},
		}, nil
	}
	if in.GetHead().GetId() == "" {
		return &pb.SaveNotesResponse{
			Head: &pb.RspHead{
				Code: -1, Msg: "req id empty",
			},
		}, nil
	}
	if in.GetNotes() == "" {
		return &pb.SaveNotesResponse{
			Head: &pb.RspHead{
				Code: -1, Msg: "req notes empty",
			},
		}, nil
	}
	err = db.SaveNotes(in.GetHead().GetId(), in.GetNotes())
	if err != nil {
		return &pb.SaveNotesResponse{
			Head: &pb.RspHead{
				Code: -1, Msg: "save notes error, " + err.Error(),
			},
		}, nil
	}
	return &pb.SaveNotesResponse{
		Head: &pb.RspHead{
			Code: 0, Msg: "success",
		},
	}, nil
}

func (s *server) GetNotes(ctx context.Context, in *pb.GetNotesRequest) (out *pb.GetNotesResponse, err error) {
	log.Printf("Received: %+v", *in)
	if in.GetHead() == nil {
		return &pb.GetNotesResponse{
			Head: &pb.RspHead{
				Code: -1, Msg: "req head empty",
			},
		}, nil
	}
	if in.GetHead().GetId() == "" {
		return &pb.GetNotesResponse{
			Head: &pb.RspHead{
				Code: -1, Msg: "req id empty",
			},
		}, nil
	}
	notes, err := db.GetNotes(in.GetHead().GetId())
	if err != nil {
		return &pb.GetNotesResponse{
			Head: &pb.RspHead{
				Code: -1, Msg: "get notes error, " + err.Error(),
			},
		}, nil
	}
	return &pb.GetNotesResponse{
		Head: &pb.RspHead{
			Code: 0, Msg: "success",
		},
		Notes: notes,
	}, nil
}
