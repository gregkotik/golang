// Package main implements a client service for parsing cvs file and sending record to Storing service.
package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    "strconv"
    "strings"

    "context"
    "time"
    "google.golang.org/grpc"
	pb "golang_test/test/record"
)

// grpc service info
const (
	address     = "localhost:50051"
	defaultName = "sender"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStoreRecordsClient(conn)

    csvFile, _ := os.Open("data.csv")
    reader := csv.NewReader(bufio.NewReader(csvFile))
    reader.Read() //skip header

    for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("-> %s\n", line)

        client_record_id, err:= strconv.ParseUint(line[0], 10, 64)
        if err != nil {
            log.Fatal(err)
        }

        // Converting to common UK format for using with the country code
        client_mobile_number := strings.TrimLeft(line[3], "(0")
        client_mobile_number = strings.Replace(client_mobile_number, ")", "", 1)
        client_mobile_number = ("+44 " + strings.Replace(client_mobile_number, " ", "", -1))

        // Contact the server and print out its response.
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()

        // Sending record
        r, err := c.StoreRecord(ctx, &pb.StoreRecordRequest{ Id : client_record_id,
                                                            Name : line[1],
                                                            Email : line[2],
                                                            MobileNumber : client_mobile_number,
                                                            })
        if err != nil {
            log.Fatalf("Could not send record: %v", err)
        }
        log.Printf("StoreRecord status: %s", r.Message)
    }
}
