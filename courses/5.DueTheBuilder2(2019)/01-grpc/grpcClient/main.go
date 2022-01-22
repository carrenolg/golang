package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"carrenolg.io/courses/dtbuilder2/01-grpc/mcache"
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
	client := mcache.NewCacheClient(conn)

	body := []byte("Hello World!")
	e := mcache.Entry{
		Key:     "foo.txt",
		Ctype:   "text/plain; charset=utf-8",
		Mtime:   time.Now().UTC().UnixNano(),
		Content: body,
		Size:    int64(len(body)),
	}
	if _, err := client.Put(context.Background(), &e); err != nil {
		log.Fatalf("Error in Put: %v", err)
	}

	ne, err := client.Get(context.Background(), &mcache.Entry{Key: "foo.txt"})
	if err != nil {
		log.Fatalf("Error in Get: %v", err)
	}
	fmt.Println(ne)
}
