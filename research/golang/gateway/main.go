package main

import (
	"context"
	"log"
	"time"

	pb "xtsgateway/gateway"

	"google.golang.org/grpc"
)

const (
	address = "xtsserver:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPushGatewayClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.OrderBook(ctx, &pb.GatewayRequest{Name: "ping"})
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	log.Printf("Response: %d", r.GetStatus())
}
