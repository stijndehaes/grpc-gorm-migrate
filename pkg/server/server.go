package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	pb "github.com/stijndehaes/gin-gonic-gorm/pb"
	"github.com/stijndehaes/gin-gonic-gorm/pkg/house"
	"github.com/stijndehaes/gin-gonic-gorm/pkg/user"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"sync"
)

type Greeter struct {
	wg sync.WaitGroup
}

// New creates new server greeter
func New() *Greeter {
	return &Greeter{}
}
func (g *Greeter) WaitStop() {
	g.wg.Wait()
}

// Start starts server
func (g *Greeter) Start() {
	log.Info("Starting")
	g.wg.Add(1)
	go func() {
		log.Fatal(g.startGRPC())
		g.wg.Done()
	}()
	g.wg.Add(1)
	go func() {
		log.Fatal(g.startREST())
		g.wg.Done()
	}()
}
func (g *Greeter) startGRPC() error {
	log.Info("Starting grpc")
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &user.Service{})
	pb.RegisterHouseServiceServer(grpcServer, &house.Service{})
	log.Info("Start serving")
	return grpcServer.Serve(lis)
}
func (g *Greeter) startREST() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:8080", opts)
	if err != nil {
		return err
	}
	err = pb.RegisterHouseServiceHandlerFromEndpoint(ctx, mux, "localhost:8080", opts)
	if err != nil {
		return err
	}
	log.Info("Start serving")
	return http.ListenAndServe("localhost:8090", mux)
}
