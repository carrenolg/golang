package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"carrenolg.io/courses/dtbuilder2/01-grpc/mcache"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type cacheServer struct {
	mu    sync.Mutex // protects store
	store map[string]*mcache.Entry
}

func (cs *cacheServer) Get(_ context.Context, e *mcache.Entry) (*mcache.Entry, error) {
	cs.mu.Lock()
	ce, ok := cs.store[e.Key]
	cs.mu.Unlock()
	if !ok {
		return e, mcache.ErrNotFound
	}
	return ce, nil
}
func (cs *cacheServer) Put(_ context.Context, e *mcache.Entry) (*mcache.Entry, error) {
	e.Stat = mcache.Status_OK
	cs.mu.Lock()
	cs.store[e.Key] = e
	cs.mu.Unlock()
	return e, nil
}
func (cs *cacheServer) Del(_ context.Context, e *mcache.Entry) (*mcache.Entry, error) {
	cs.mu.Lock()
	delete(cs.store, e.Key)
	cs.mu.Unlock()
	return e, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// Lanzamos el server gRPC
	srv := grpc.NewServer()

	// Atamos nuestro server a la interface del paquete mcache generada por protoc.
	mcache.RegisterCacheServer(srv, &cacheServer{store: make(map[string]*mcache.Entry)})

	// Iniciamos el server.
	fmt.Printf("Servidor mcache gRPC en puerto %s listo. CTRL+C para detener.\n", port)
	if err := srv.Serve(lis); err != nil && err != grpc.ErrServerStopped {
		// Error iniciando el Server. Posible conflicto de puerto, permisos, etc.
		log.Printf("Error durante Serve: %v", err)
	}

}
