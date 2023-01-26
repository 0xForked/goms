package grpc

import (
	"context"
	"errors"
	"github.com/bakode/goms/internal/store/domain/entity"
	"github.com/bakode/goms/mocks"
	"github.com/bakode/goms/pkg/pb"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
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

type storeGRPCHandlerTestSuite struct {
	suite.Suite
	svc mocks.IStoreRepository
}

func GetConnWithMock(s *storeGRPCHandlerTestSuite, mockType MockType) *grpc.ClientConn {
	lsn := bufconn.Listen(1024 * 1024)
	s.Suite.T().Cleanup(func() { _ = lsn.Close() })

	svr := grpc.NewServer()
	s.Suite.T().Cleanup(func() { svr.Stop() })

	svc := new(mocks.IStoreService)
	handler := NewStoreGRPCHandler(svc)
	pb.RegisterStoreGRPCHandlerServer(svr, handler)
	go func() {
		if err := svr.Serve(lsn); err != nil {
			log.Fatalf("svr.Serve %v", err.Error())
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lsn.Dial()
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

	switch mockType {
	case MockFetchSuccess:
		svc.On("All", mock.Anything).Return([]*entity.Store{
			{ID: 1, Name: "lorem"},
			{ID: 2, Name: "ipsum"},
		}, nil)
	case MockFetchError:
		svc.On("All", mock.Anything).
			Return(nil, errors.New("lorem"))
	case MockShowSuccess:
		svc.On("Find", mock.Anything, mock.Anything).
			Return(&entity.Store{ID: 1, Name: "lorem"}, nil)
	case MockShowError:
		svc.On("Find", mock.Anything, mock.Anything).
			Return(nil, errors.New("lorem"))
	case MockStoreSuccess:
		svc.On("Record", mock.Anything, mock.Anything).
			Return(nil)
	case MockStoreError:
		svc.On("Record", mock.Anything, mock.Anything).
			Return(errors.New("lorem"))
	case MockUpdateSuccess:
		svc.On("Patch", mock.Anything, mock.Anything).
			Return(nil)
	case MockUpdateError:
		svc.On("Patch", mock.Anything, mock.Anything).
			Return(errors.New("lorem"))
	case MockDestroySuccess:
		svc.On("Erase", mock.Anything, mock.Anything).
			Return(nil)
	case MockDestroyError:
		svc.On("Erase", mock.Anything, mock.Anything).
			Return(errors.New("lorem"))
	}

	return conn
}

// FETCH
func (suite *storeGRPCHandlerTestSuite) TestHandler_Fetch_ShouldSuccess() {
	conn := GetConnWithMock(suite, MockFetchSuccess)
	client := pb.NewStoreGRPCHandlerClient(conn)
	res, err := client.Fetch(context.Background(), &pb.StoreEmptyRequest{})
	suite.Nil(err)
	suite.NotNil(res)
	suite.Equal(len(res.Stores), 2)
}
func (suite *storeGRPCHandlerTestSuite) TestHandler_Fetch_ShouldError() {
	conn := GetConnWithMock(suite, MockFetchError)
	client := pb.NewStoreGRPCHandlerClient(conn)
	res, err := client.Fetch(context.Background(), &pb.StoreEmptyRequest{})
	suite.Nil(res)
	suite.NotNil(err)
	suite.Contains(err.Error(), "lorem")
}

// SHOW
func (suite *storeGRPCHandlerTestSuite) TestHandler_Show_ShouldSuccess() {
	conn := GetConnWithMock(suite, MockShowSuccess)
	client := pb.NewStoreGRPCHandlerClient(conn)
	res, err := client.Show(context.Background(), &pb.StoreIDModel{Id: 1})
	suite.Nil(err)
	suite.NotNil(res)
	suite.Equal(res.Store.Name, "lorem")
}
func (suite *storeGRPCHandlerTestSuite) TestHandler_Show_ShouldError() {
	conn := GetConnWithMock(suite, MockShowError)
	client := pb.NewStoreGRPCHandlerClient(conn)
	res, err := client.Show(context.Background(), &pb.StoreIDModel{Id: 1})
	suite.Nil(res)
	suite.NotNil(err)
	suite.Contains(err.Error(), "lorem")
}

// STORE
func (suite *storeGRPCHandlerTestSuite) TestHandler_Store_ShouldSuccess() {
	conn := GetConnWithMock(suite, MockStoreSuccess)
	client := pb.NewStoreGRPCHandlerClient(conn)
	res, err := client.Store(context.Background(), &pb.StoreNameModel{Name: "lorem"})
	suite.Nil(err)
	suite.NotNil(res)
	suite.Equal(res.Status, true)
}
func (suite *storeGRPCHandlerTestSuite) TestHandler_Store_ShouldError() {
	conn := GetConnWithMock(suite, MockStoreError)
	client := pb.NewStoreGRPCHandlerClient(conn)
	res, err := client.Store(context.Background(), &pb.StoreNameModel{Name: "lorem"})
	suite.NotNil(err)
	suite.Nil(res)
	suite.Contains(err.Error(), "lorem")
}

// UPDATE
func (suite *storeGRPCHandlerTestSuite) TestHandler_Update_ShouldSuccess() {
	conn := GetConnWithMock(suite, MockUpdateSuccess)
	client := pb.NewStoreGRPCHandlerClient(conn)
	res, err := client.Update(context.Background(), &pb.StoreModel{Id: 1, Name: "lorem"})
	suite.Nil(err)
	suite.NotNil(res)
	suite.Equal(res.Status, true)
}
func (suite *storeGRPCHandlerTestSuite) TestHandler_Update_ShouldError() {
	conn := GetConnWithMock(suite, MockUpdateError)
	client := pb.NewStoreGRPCHandlerClient(conn)
	res, err := client.Update(context.Background(), &pb.StoreModel{Id: 1, Name: "lorem"})
	suite.NotNil(err)
	suite.Nil(res)
	suite.Contains(err.Error(), "lorem")
}

// DESTROY
func (suite *storeGRPCHandlerTestSuite) TestHandler_Destroy_ShouldSuccess() {
	conn := GetConnWithMock(suite, MockDestroySuccess)
	client := pb.NewStoreGRPCHandlerClient(conn)
	res, err := client.Destroy(context.Background(), &pb.StoreIDModel{Id: 1})
	suite.Nil(err)
	suite.NotNil(res)
	suite.Equal(res.Status, true)
}
func (suite *storeGRPCHandlerTestSuite) TestHandler_Destroy_ShouldError() {
	conn := GetConnWithMock(suite, MockDestroyError)
	client := pb.NewStoreGRPCHandlerClient(conn)
	res, err := client.Destroy(context.Background(), &pb.StoreIDModel{Id: 1})
	suite.NotNil(err)
	suite.Nil(res)
	suite.Contains(err.Error(), "lorem")
}

func TestStoreGRPCHandlerService(t *testing.T) {
	suite.Run(t, new(storeGRPCHandlerTestSuite))
}
