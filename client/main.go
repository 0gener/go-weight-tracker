package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/0gener/go-weight-tracker/weighttracker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	// createRecord(c)
	readRecord(c, 1)
}

func createRecord(c weighttracker.WeightTrackerClient) {
	fmt.Println("calling CreateRecord")

	record := &weighttracker.Record{
		Weight:     63.3,
		WeightedAt: timestamppb.New(time.Date(2019, 10, 10, 9, 12, 56, 100, time.Local)),
	}

	res, err := c.CreateRecord(context.Background(), &weighttracker.CreateRecordRequest{Record: record})
	if err != nil {
		log.Printf("failed to AddRecord: %v\n", err)
	}

	fmt.Printf("CreateRecord result: %v\n", res)
}

func readRecord(c weighttracker.WeightTrackerClient, recordID uint64) {
	fmt.Println("calling ReadRecord")

	res, err := c.ReadRecord(context.Background(), &weighttracker.ReadRecordRequest{RecordId: recordID})
	if err != nil {
		log.Printf("failed to ReadRecord: %v\n", err)
	}

	fmt.Printf("ReadRecord result: %v\n", res)
}
