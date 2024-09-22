package app

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/erry-az/otel-go"
	"github.com/erry-az/otel-go/example/otelgrpc/api"
	handlerGrpc "github.com/erry-az/otel-go/example/otelgrpc/server/handler/grpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	defaultAddress = "0.0.0.0:7777"
)

// Grpc struct that hold grpc app server requirement
type Grpc struct {
	otel       *otel.Providers
	grpcServer *grpc.Server

	grpcAddress string
	ctx         context.Context
}

// NewGrpc initialize grpc app
func NewGrpc(ctx context.Context) (*Grpc, error) {
	grpcAddress := os.Getenv("ADDRESS")
	if grpcAddress == "" {
		grpcAddress = defaultAddress
	}

	// init open telemetry
	otelProviders, err := otel.NewProviders(ctx)
	if err != nil {
		return nil, err
	}

	// init grpc Grpc
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	reflection.Register(grpcServer)

	api.RegisterHelloServiceServer(grpcServer, &handlerGrpc.HelloService{})

	return &Grpc{
		otel:        otelProviders,
		grpcServer:  grpcServer,
		grpcAddress: grpcAddress,
		ctx:         ctx,
	}, err
}

// Start run grpc app
func (s *Grpc) Start() error {
	lis, err := net.Listen("tcp", s.grpcAddress)
	if err != nil {
		return err
	}

	log.Printf("grpc start at: %s", s.grpcAddress)

	return s.grpcServer.Serve(lis)
}

// Stop shutdown or stop all grpc app requirement
func (s *Grpc) Stop() error {
	timeoutCtx, cancel := context.WithTimeout(s.ctx, time.Second*5)
	defer cancel()

	s.grpcServer.GracefulStop()

	return s.otel.Shutdown(timeoutCtx)
}
