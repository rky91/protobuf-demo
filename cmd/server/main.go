package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "protobuf-demo/pb/proto"
	"protobuf-demo/service"
)

func main() {

	port := flag.Int("port", 0, "server port")
	flag.Parse()
	fmt.Printf("The server port is : %d \n", *port)

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Errorf("unable to start the server : %s", address)
	}

	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Errorf("unable to start the server : %s", address)
	}
}
