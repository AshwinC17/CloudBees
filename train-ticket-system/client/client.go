package main

import (
	"context"
	"log"

	"github.com/AshwinC17/train-ticket-system/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewTicketServiceClient(conn)

	req := &proto.TicketRequest{
		From:  "Station A",
		To:    "Station B",
		User:  &proto.User{Email: "test@example.com"},
		Price: 100,
	}

	resp, err := client.PurchaseTicket(context.Background(), req)
	if err != nil {
		log.Fatalf("could not purchase ticket: %v", err)
	}

	log.Printf("Purchased ticket: %v", resp)
}
