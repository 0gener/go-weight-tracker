package main

import (
	"context"
	"fmt"
	"log"

	"github.com/0gener/go-weight-tracker/weighttracker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls      = false
	certFile = "ssl/ca.crt"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	opts := grpc.WithInsecure()

	if tls {
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")

		if sslErr != nil {
			log.Fatalf("error while loading CA trust certificate: %v", sslErr)
		}

		opts = grpc.WithTransportCredentials(creds)
	}

	conn, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := weighttracker.NewWeightTrackerClient(conn)

	record := &weighttracker.Record{
		Weight: 68,
	}

	res, err := c.AddRecord(context.Background(), &weighttracker.AddRecordRequest{Record: record})
	if err != nil {
		log.Printf("failed to AddRecord: %v\n", err)
	}

	fmt.Println(res)
}
