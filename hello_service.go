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
