package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	pb "protobuf-demo/pb/proto"
	"protobuf-demo/sample"
)

func main() {

	serverAddress := flag.String("address", "", "server address")
	flag.Parse()

	fmt.Printf("The server address is : %s \n", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not dial the server : %s", *serverAddress)
	}

	client := pb.NewLaptopServiceClient(conn)
	laptop := sample.NewLaptop()
	laptop.Id = "5a03195f-7364-433b-8785-9fa8b195274a"
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	laptopRes, err := client.CreateLaptop(context.Background(), req)

	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("Laptop already exists.")
		} else {

		}
		log.Fatal("Unable to create the laptop :", err)
		return
	}

	log.Printf("Laptop created successfully : %s", laptopRes.Id)
}
