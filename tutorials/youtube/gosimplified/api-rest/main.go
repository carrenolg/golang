package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"carrenolg.io/tutorial/gosimplifies/api-rest/server"
)

func main() {
	// main
	ctx := context.Background()

	serverDone := make(chan os.Signal, 1)
	signal.Notify(serverDone, os.Interrupt, syscall.SIGTERM)

	srv := server.New(":8080")

	go srv.ListenAndServe()
	log.Println("server started")

	<-serverDone
	srv.Shutdown(ctx)
	log.Println("server stopped")
}
