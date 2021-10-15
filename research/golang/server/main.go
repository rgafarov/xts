package main

import (
	"context"
	"log"
	"net"

	pb "xtsserver/gateway"

	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedPushGatewayServer
}

const (
	port = ":50051"
)

func (s *server) OrderBook(ctx context.Context, in *pb.GatewayRequest) (*pb.GatewayResponse, error) {
	return &pb.GatewayResponse{
		Status: 7,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPushGatewayServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
