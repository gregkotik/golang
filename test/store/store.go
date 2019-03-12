// Package main implements a server for Storing record service.
package main

import (
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	pb "golang_test/test/record"
)

const (
	port = ":50051"
)

type DbRecord struct  {
    id uint64
    name string
    email string
    mobileNumber string
}
// Delcaring map
var m map[uint64]*DbRecord

// Server is used to implement Record Storing Service
type server struct{}

// StoreRecord implements storerecords.StoreRecordServer
func (s *server) StoreRecord(ctx context.Context, in *pb.StoreRecordRequest) (*pb.StoreRecordReply, error) {
	log.Printf("Received: %d %v %v %v", in.Id, in.Name, in.Email, in.MobileNumber)

	// Adds new element to the map.
	// If element exists, overwrites the existed element with new values
    m[in.Id] = &DbRecord{in.Id, in.Name, in.Email, in.MobileNumber}

	return &pb.StoreRecordReply{Message: "Stored " + strconv.FormatUint(in.Id, 10) + " " + in.Name}, nil
}

func main() {
    // Creating map
    m = make(map[uint64]*DbRecord)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStoreRecordsServer(s, &server{})

    // Start server loop
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
