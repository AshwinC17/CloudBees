package main

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"

	"github.com/AshwinC17/train-ticket-system/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedTicketServiceServer
	mu      sync.Mutex
	tickets map[string]*proto.TicketResponse
	seats   map[string]string
}

func (s *server) PurchaseTicket(ctx context.Context, req *proto.TicketRequest) (*proto.TicketResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ticket := &proto.TicketResponse{
		From:  req.From,
		To:    req.To,
		User:  req.User,
		Price: req.Price,
		Seat:  "A1", // This should be dynamically allocated
	}

	s.tickets[req.User.Email] = ticket
	s.seats[req.User.Email] = "A" // This should be dynamically allocated

	return ticket, nil
}

func (s *server) GetReceipt(ctx context.Context, user *proto.User) (*proto.TicketResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ticket, ok := s.tickets[user.Email]
	if !ok {
		return nil, errors.New("no ticket found for user")
	}

	return ticket, nil
}

func (s *server) GetUsersBySection(ctx context.Context, req *proto.SectionRequest) (*proto.UserSeatResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	response := &proto.UserSeatResponse{}

	for email, section := range s.seats {
		if section == req.Section {
			user := &proto.User{Email: email}
			response.Users = append(response.Users, user)
		}
	}

	return response, nil
}

func (s *server) RemoveUser(ctx context.Context, user *proto.User) (*proto.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.tickets, user.Email)
	delete(s.seats, user.Email)

	return &proto.Empty{}, nil
}

func (s *server) ModifySeat(ctx context.Context, req *proto.ModifySeatRequest) (*proto.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.seats[req.User.Email] = req.NewSection

	return &proto.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterTicketServiceServer(s, &server{
		tickets: make(map[string]*proto.TicketResponse),
		seats:   make(map[string]string),
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
