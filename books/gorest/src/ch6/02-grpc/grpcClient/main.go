package main

import (
	"context"
	"log"

	pb "carrenolg.io/books/gorest/src/ch6/02-grpc/datafiles"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMoneyTransactionClient(conn)

	// Prepate data. Get this from clients like Fronted or App
	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	// Contact the server and print out its response.
	r, err := c.MakeTransaction(context.Background(), &pb.TransactionRequest{
		From:   from,
		To:     to,
		Amount: amount,
	})
	if err != nil {
		log.Fatalf("Could not transact: %v", err)
	}
	log.Printf("Trasaction confirmed: %t", r.Confirmation)
}
