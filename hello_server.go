package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"

	. "github.com/NFCHKK/grpc/proto"
	"google.golang.org/grpc"
)

var (
	port               = flag.Int("port", 9091, "server port")
	grpcServerEndPoint = flag.String("grpc-server-endpoint",
		"localhost:9091",
		"grpc server endpoint")
)

type server struct {
	UnimplementedHelloServiceServer
}

// for test

// curl -X POST -k http://localhost:8081/example/echo -d '{"name": " hello"}'
func run() error {
	// Start a http server and a rpc client
	// use rpc client connect to rpc server
	// translate http request to rpc request then translate rpc response to http response
	// 3 different ways to gen stub file
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := RegisterHelloServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndPoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9091))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterHelloServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	// run rpc server
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to server: %v", err)
		}
	}()
	// run gateway proxy: start http server and connect rpc server
	if err := run(); err != nil {
		log.Fatalf("failed to start gate server: %v", err)
	}
}
