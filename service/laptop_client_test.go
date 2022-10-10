package service_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"net"
	"protobuf-demo/pb/proto"
	"protobuf-demo/sample"
	"protobuf-demo/service"
	"testing"
)

func startTestLaptopServer(t *testing.T, laptopStore service.LaptopStore) string {
	laptopServer := service.NewLaptopServer(laptopStore)

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0") // random available port
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopStore := service.NewInMemoryLaptopStore()
	serverAddress := startTestLaptopServer(t, laptopStore)
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedID := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	// check that the laptop is saved to the store
	other, err := laptopStore.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	// check that the saved laptop is the same as the one we send
	require.Equal(t, laptop.Name, other.Name)
}
