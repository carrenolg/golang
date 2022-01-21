package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "carrenolg.io/books/gorest/src/ch6/02-grpc/datafiles"
)

const (
	port = ":50051"
)

// server is used to create MoneyTransactionServer
type server struct{}

// MakeTransaction implements MoneyTransactionServer.MakeTransaction
func (s *server) MakeTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	log.Println("Got request for money Transfer...")
	log.Printf("Amount: %f, From A/c: %s, To A/c: %s", in.Amount, in.From, in.To)
	// Do database logic here ....
	return &pb.TransactionResponse{
		Confirmation: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	srv := &server{}
	pb.RegisterMoneyTransactionServer(s, srv)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
