package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/0gener/go-weight-tracker/weighttracker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	host     = "0.0.0.0"
	port     = 50051
	tls      = false
	certFile = "" // required if tls enabled
	keyFile  = "" // required if tls enabled
)

type server struct {
	weighttracker.UnsafeWeightTrackerServer
}

func (*server) AddRecord(ctx context.Context, req *weighttracker.AddRecordRequest) (*weighttracker.AddRecordResponse, error) {
	log.Printf("called AddRecord: %v\n", req)

	return &weighttracker.AddRecordResponse{
		Record: req.GetRecord(),
	}, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("starting server on port %d...\n", port)

	opts := []grpc.ServerOption{}

	if tls {
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)

		if sslErr != nil {
			log.Fatalf("failed to load certificates: %v\n", sslErr)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	sv := grpc.NewServer(opts...)

	weighttracker.RegisterWeightTrackerServer(sv, &server{})

	go func() {
		if err = sv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v\n", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

	log.Println("stopping server...")

	sv.Stop()
	lis.Close()
}
