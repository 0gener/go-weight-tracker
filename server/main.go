package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/0gener/go-weight-tracker/weighttracker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     = "0.0.0.0"
	port     = 50051
	tls      = false
	certFile = "" // required if tls enabled
	keyFile  = "" // required if tls enabled
)

var (
	mysqlHost     = "localhost"
	mysqlPort     = 3306
	mysqlSchema   = "weight_tracker"
	mysqlUser     = "root"
	mysqlPassword = "123456"
)

var db2 *gorm.DB

type server struct {
	weighttracker.UnsafeWeightTrackerServer
}

// Record is ...
type Record struct {
	gorm.Model
	Weight     float32   `gorm:"type:decimal(4,2);not null"`
	WeightedAt time.Time `gorm:"not null"`
}

func (*server) CreateRecord(ctx context.Context, req *weighttracker.CreateRecordRequest) (*weighttracker.CreateRecordResponse, error) {
	log.Printf("called CreateRecord: %v\n", req)

	if req.GetRecord().GetWeight() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "weight must be greater than 0")
	}

	var recordDatetime time.Time
	if req.GetRecord().GetWeightedAt() != nil {
		recordDatetime = req.GetRecord().GetWeightedAt().AsTime()
	} else {
		recordDatetime = time.Now()
	}

	record := Record{
		Weight:     req.GetRecord().GetWeight(),
		WeightedAt: recordDatetime,
	}

	res := db2.Create(&record)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error while inserting record on db: %v", res.Error))
	}

	return &weighttracker.CreateRecordResponse{
		Record: dataToRecordPb(record),
	}, nil
}

func (*server) ReadRecord(ctx context.Context, req *weighttracker.ReadRecordRequest) (*weighttracker.ReadRecordResponse, error) {
	log.Printf("called ReadRecord: %v\n", req)

	recordID := req.GetRecordId()

	record := Record{}
	res := db2.First(&record, recordID)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("no record found with id = %d", record.ID))
		}

		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error while reading record from db: %v", res.Error))
	}

	return &weighttracker.ReadRecordResponse{
		Record: dataToRecordPb(record),
	}, nil
}

func (*server) UpdateRecord(ctx context.Context, req *weighttracker.UpdateRecordRequest) (*weighttracker.UpdateRecordResponse, error) {
	log.Printf("called UpdateRecord: %v\n", req)

	var recordDatetime time.Time
	if req.GetRecord().GetWeightedAt() != nil {
		recordDatetime = req.GetRecord().GetWeightedAt().AsTime()
	}

	record := Record{}
	res := db2.Model(&record).Where("id = ?", req.GetRecord().GetId()).Updates(Record{
		Weight:     req.GetRecord().GetWeight(),
		WeightedAt: recordDatetime,
	})
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("no record found with id = %d", req.GetRecord().GetId()))
		}

		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error while updating record from db: %v", res.Error))
	}

	return &weighttracker.UpdateRecordResponse{
		Record: dataToRecordPb(record),
	}, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	connectMySQL()

	startServer()
}

func connectMySQL() {
	log.Println("connecting to mysql...")

	dbURL := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?charset=utf8mb4&parseTime=True", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlSchema)

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed connect to mysql: %v\n", err)
	}

	db2 = db

	db2.AutoMigrate(&Record{})
}

func startServer() {
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

func dataToRecordPb(rec Record) *weighttracker.Record {
	return &weighttracker.Record{
		Id:         uint64(rec.ID),
		Weight:     rec.Weight,
		WeightedAt: timestamppb.New(rec.WeightedAt),
	}
}
