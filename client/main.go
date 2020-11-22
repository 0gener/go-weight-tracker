package main

import (
	"context"
	"fmt"
	"io"
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
	// readRecord(c, 2)
	// updateRecord(c, 1)
	// deleteRecord(c, 1)
	listRecords(c)
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

func updateRecord(c weighttracker.WeightTrackerClient, recordID uint64) {
	fmt.Println("calling UpdateRecord")

	res, err := c.UpdateRecord(context.Background(), &weighttracker.UpdateRecordRequest{
		Record: &weighttracker.Record{
			Id:     recordID,
			Weight: 63,
		},
	})
	if err != nil {
		log.Printf("failed to UpdateRecord: %v\n", err)
	}

	fmt.Printf("UpdateRecord result: %v\n", res)
}

func deleteRecord(c weighttracker.WeightTrackerClient, recordID uint64) {
	fmt.Println("calling DeleteRecord")

	res, err := c.DeleteRecord(context.Background(), &weighttracker.DeleteRecordRequest{RecordId: recordID})
	if err != nil {
		log.Printf("failed to DeleteRecord: %v\n", err)
	}

	fmt.Printf("DeleteRecord result: %v\n", res)
}

func listRecords(c weighttracker.WeightTrackerClient) {
	fmt.Println("calling ListRecord")

	// from, _ := time.Parse("2006-01-02", "2020-10-10")

	stream, err := c.ListRecords(context.Background(), &weighttracker.ListRecordsRequest{
		// WeightedAtFrom: timestamppb.New(from),
		WeightedAtTo: timestamppb.Now(),
	})
	if err != nil {
		log.Printf("failed to ListRecords: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("something happened: %v", err)
		}

		fmt.Println(res.GetRecord())
	}
}
