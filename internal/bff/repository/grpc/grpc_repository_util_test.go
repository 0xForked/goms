package grpc_test

import (
	"context"
	"errors"
	"github.com/aasumitro/goms/internal/bff/domain/contract"
	grpcRepo "github.com/aasumitro/goms/internal/bff/repository/grpc"
	"github.com/aasumitro/goms/mocks"
	"github.com/aasumitro/goms/pkg/pb"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
)

type MockType int64

const (
	MockFetchSuccess MockType = iota
	MockFetchError            = iota
	MockShowSuccess
	MockShowError
	MockStoreSuccess
	MockStoreError
	MockUpdateSuccess
	MockUpdateError
	MockDestroySuccess
	MockDestroyError
)

// TODO: MAKE IT REUSABLE
func mockBookGRPCConnection(s *bookGRPCRepositoryTestSuite, mockType MockType, name string) contract.IBookGRPCRepository {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	serverMock := new(mocks.BookGRPCHandlerServer)
	pb.RegisterBookGRPCHandlerServer(server, serverMock)
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	dialer := func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
	conn, err := grpc.DialContext(
		context.Background(), "",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(
			insecure.NewCredentials()))
	if err != nil {
		s.Suite.T().Fatalf("grpc.DialContext %v", err.Error())
	}
	s.Suite.T().Cleanup(func() { _ = conn.Close() })
	client := pb.NewBookGRPCHandlerClient(conn)
	switch mockType {
	case MockFetchSuccess:
		serverMock.On("Fetch", mock.Anything, mock.Anything).
			Return(&pb.BookRowsResponse{
				Books: []*pb.BookModel{
					{Id: 1, StoreId: 1, Name: "lorem"},
					{Id: 2, StoreId: 1, Name: "ipsum"},
				},
			}, nil)
	case MockFetchError:
		serverMock.On("Fetch", mock.Anything, mock.Anything).
			Return(nil, errors.New("lorem"))
	case MockShowSuccess:
		serverMock.On("Show", mock.Anything, mock.Anything).
			Return(&pb.BookRowResponse{
				Book: &pb.BookModel{Id: 1, StoreId: 1, Name: "lorem"},
			}, nil)
	case MockShowError:
		serverMock.On("Show", mock.Anything, mock.Anything).
			Return(nil, errors.New("lorem"))
	case MockStoreSuccess, MockUpdateSuccess, MockDestroySuccess:
		serverMock.On(name, mock.Anything, mock.Anything).
			Return(&pb.BookBoolResponse{Status: true}, nil)
	case MockDestroyError, MockUpdateError, MockStoreError:
		serverMock.On(name, mock.Anything, mock.Anything).
			Return(&pb.BookBoolResponse{Status: false}, errors.New("lorem"))
	}
	return grpcRepo.NewBookGRPCRepository(client)
}

// TODO: MAKE IT REUSABLE
func mockStoreGRPCConnection(s *storeGRPCRepositoryTestSuite, mockType MockType, name string) contract.IStoreGRPCRepository {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	serverMock := new(mocks.StoreGRPCHandlerServer)
	pb.RegisterStoreGRPCHandlerServer(server, serverMock)
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	dialer := func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
	conn, err := grpc.DialContext(
		context.Background(), "",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(
			insecure.NewCredentials()))
	if err != nil {
		s.Suite.T().Fatalf("grpc.DialContext %v", err.Error())
	}
	s.Suite.T().Cleanup(func() { _ = conn.Close() })
	client := pb.NewStoreGRPCHandlerClient(conn)
	switch mockType {
	case MockFetchSuccess:
		serverMock.On("Fetch", mock.Anything, mock.Anything).
			Return(&pb.StoreRowsResponse{
				Stores: []*pb.StoreModel{
					{Id: 1, Name: "lorem"},
					{Id: 2, Name: "ipsum"},
				},
			}, nil)
	case MockFetchError:
		serverMock.On("Fetch", mock.Anything, mock.Anything).
			Return(nil, errors.New("lorem"))
	case MockShowSuccess:
		serverMock.On("Show", mock.Anything, mock.Anything).
			Return(&pb.StoreRowResponse{
				Store: &pb.StoreModel{Id: 1, Name: "lorem"},
			}, nil)
	case MockShowError:
		serverMock.On("Show", mock.Anything, mock.Anything).
			Return(nil, errors.New("lorem"))
	case MockStoreSuccess, MockUpdateSuccess, MockDestroySuccess:
		serverMock.On(name, mock.Anything, mock.Anything).
			Return(&pb.StoreBoolResponse{Status: true}, nil)
	case MockDestroyError, MockUpdateError, MockStoreError:
		serverMock.On(name, mock.Anything, mock.Anything).
			Return(&pb.StoreBoolResponse{Status: false}, errors.New("lorem"))
	}
	return grpcRepo.NewStoreGRPCRepository(client)
}
