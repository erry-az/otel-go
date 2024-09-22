package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/erry-az/otel-go/example/otelgrpc/server/app"
)

func main() {
	capturedSignal := make(chan os.Signal, 1)
	signal.Notify(capturedSignal, syscall.SIGINT, syscall.SIGTERM)

	myServer, err := app.NewGrpc(context.Background())
	if err != nil {
		log.Fatalf("failed create server: %v", err)
	}

	go func() {
		err := myServer.Start()
		if err != nil {
			log.Fatalf("failed start server: %v", err)
		}
	}()

	<-capturedSignal

	err = myServer.Stop()
	if err != nil {
		log.Fatalf("failed shutdown server: %v", err)
	}
}
