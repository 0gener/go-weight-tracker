package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/0gener/go-weight-tracker/weighttracker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

func (*server) AddRecord(ctx context.Context, req *weighttracker.AddRecordRequest) (*weighttracker.AddRecordResponse, error) {
	log.Printf("called AddRecord: %v\n", req)

	var recordDatetime time.Time
	if req.Record.WeightedAt != nil {
		recordDatetime = req.GetRecord().WeightedAt.AsTime()
	} else {
		recordDatetime = time.Now()
	}

	record := Record{
		Weight:     req.GetRecord().GetWeight(),
		WeightedAt: recordDatetime,
	}

	db2.Create(&record)

	return &weighttracker.AddRecordResponse{
		Record: &weighttracker.Record{
			Id:         uint64(record.ID),
			Weight:     record.Weight,
			WeightedAt: timestamppb.New(record.WeightedAt),
		},
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
