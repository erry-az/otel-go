package grpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/erry-az/otel-go/example/otelgrpc/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"log"
	"time"
)

var tracer = otel.Tracer("grpc-example-server")

// HelloService is used to implement api.HelloServiceServer.
type HelloService struct {
	api.UnimplementedHelloServiceServer
}

// SayHello implements api.HelloServiceServer.
func (s *HelloService) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	log.Printf("Received: %v\n", in.GetGreeting())
	s.workHard(ctx)
	time.Sleep(50 * time.Millisecond)

	return &api.HelloResponse{Reply: "Hello " + in.Greeting}, nil
}

func (s *HelloService) workHard(ctx context.Context) {
	_, span := tracer.Start(ctx, "workHard",
		trace.WithAttributes(attribute.String("extra.key", "extra.value")))
	defer span.End()

	time.Sleep(50 * time.Millisecond)
}

func (s *HelloService) SayHelloServerStream(in *api.HelloRequest, out api.HelloService_SayHelloServerStreamServer) error {
	log.Printf("Received: %v\n", in.GetGreeting())

	for i := 0; i < 5; i++ {
		err := out.Send(&api.HelloResponse{Reply: "Hello " + in.Greeting})
		if err != nil {
			return err
		}

		time.Sleep(time.Duration(i*50) * time.Millisecond)
	}

	return nil
}

func (s *HelloService) SayHelloClientStream(stream api.HelloService_SayHelloClientStreamServer) error {
	i := 0

	for {
		in, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Printf("Non EOF error: %v\n", err)
			return err
		}

		log.Printf("Received: %v\n", in.GetGreeting())
		i++
	}

	time.Sleep(50 * time.Millisecond)

	return stream.SendAndClose(&api.HelloResponse{Reply: fmt.Sprintf("Hello (%v times)", i)})
}

func (s *HelloService) SayHelloBidiStream(stream api.HelloService_SayHelloBidiStreamServer) error {
	for {
		in, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Printf("Non EOF error: %v\n", err)
			return err
		}

		time.Sleep(50 * time.Millisecond)

		log.Printf("Received: %v\n", in.GetGreeting())
		err = stream.Send(&api.HelloResponse{Reply: "Hello " + in.Greeting})
		if err != nil {
			return err
		}
	}

	return nil
}
