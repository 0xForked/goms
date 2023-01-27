package grpc_test

import (
	"context"
	"errors"
	delivery "github.com/aasumitro/goms/internal/book/delivery/handler/grpc"
	"github.com/aasumitro/goms/internal/book/domain/entity"
	"github.com/aasumitro/goms/mocks"
	"github.com/aasumitro/goms/pkg/pb"
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

type bookGRPCHandlerTestSuite struct {
	suite.Suite
}

func GetConnWithMock(s *bookGRPCHandlerTestSuite, mockType MockType) *grpc.ClientConn {
	lsn := bufconn.Listen(1024 * 1024)
	s.Suite.T().Cleanup(func() { _ = lsn.Close() })

	svr := grpc.NewServer()
	s.Suite.T().Cleanup(func() { svr.Stop() })

	svc := new(mocks.IBookService)
	handler := delivery.NewBookGRPCHandler(svc)
	pb.RegisterBookGRPCHandlerServer(svr, handler)
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
		svc.On("All", mock.Anything).Return([]*entity.Book{
			{ID: 1, StoreID: 1, Name: "lorem"},
			{ID: 2, StoreID: 1, Name: "ipsum"},
		}, nil)
	case MockFetchError:
		svc.On("All", mock.Anything, mock.Anything).
			Return(nil, errors.New("lorem"))
	case MockShowSuccess:
		svc.On("Find", mock.Anything, mock.Anything).
			Return(&entity.Book{ID: 1, StoreID: 1, Name: "lorem"}, nil)
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
func (suite *bookGRPCHandlerTestSuite) TestHandler_Fetch_ShouldSuccess() {
	conn := GetConnWithMock(suite, MockFetchSuccess)
	client := pb.NewBookGRPCHandlerClient(conn)
	res, err := client.Fetch(context.Background(), &pb.BookIDModel{})
	suite.Nil(err)
	suite.NotNil(res)
	suite.Equal(len(res.Books), 2)
}
func (suite *bookGRPCHandlerTestSuite) TestHandler_Fetch_ShouldError() {
	conn := GetConnWithMock(suite, MockFetchError)
	client := pb.NewBookGRPCHandlerClient(conn)
	res, err := client.Fetch(context.Background(), &pb.BookIDModel{Type: pb.ActionType_RELATED, Id: 1})
	suite.Nil(res)
	suite.NotNil(err)
	suite.Contains(err.Error(), "lorem")
}

// SHOW
func (suite *bookGRPCHandlerTestSuite) TestHandler_Show_ShouldSuccess() {
	conn := GetConnWithMock(suite, MockShowSuccess)
	client := pb.NewBookGRPCHandlerClient(conn)
	res, err := client.Show(context.Background(), &pb.BookIDModel{Type: pb.ActionType_SPECIFIED, Id: 1})
	suite.Nil(err)
	suite.NotNil(res)
	suite.Equal(res.Book.Name, "lorem")
}
func (suite *bookGRPCHandlerTestSuite) TestHandler_Show_ShouldError() {
	conn := GetConnWithMock(suite, MockShowError)
	client := pb.NewBookGRPCHandlerClient(conn)
	res, err := client.Show(context.Background(), &pb.BookIDModel{Type: pb.ActionType_SPECIFIED, Id: 1})
	suite.Nil(res)
	suite.NotNil(err)
	suite.Contains(err.Error(), "lorem")
}

// STORE
func (suite *bookGRPCHandlerTestSuite) TestHandler_Store_ShouldSuccess() {
	conn := GetConnWithMock(suite, MockStoreSuccess)
	client := pb.NewBookGRPCHandlerClient(conn)
	res, err := client.Store(context.Background(), &pb.BookAddRequest{Name: "lorem", StoreId: 1})
	suite.Nil(err)
	suite.NotNil(res)
	suite.Equal(res.Status, true)
}
func (suite *bookGRPCHandlerTestSuite) TestHandler_Store_ShouldError() {
	conn := GetConnWithMock(suite, MockStoreError)
	client := pb.NewBookGRPCHandlerClient(conn)
	res, err := client.Store(context.Background(), &pb.BookAddRequest{Name: "lorem", StoreId: 1})
	suite.NotNil(err)
	suite.Nil(res)
	suite.Contains(err.Error(), "lorem")
}

// UPDATE
func (suite *bookGRPCHandlerTestSuite) TestHandler_Update_ShouldSuccess() {
	conn := GetConnWithMock(suite, MockUpdateSuccess)
	client := pb.NewBookGRPCHandlerClient(conn)
	res, err := client.Update(context.Background(), &pb.BookModel{Id: 1, StoreId: 1, Name: "lorem"})
	suite.Nil(err)
	suite.NotNil(res)
	suite.Equal(res.Status, true)
}
func (suite *bookGRPCHandlerTestSuite) TestHandler_Update_ShouldError() {
	conn := GetConnWithMock(suite, MockUpdateError)
	client := pb.NewBookGRPCHandlerClient(conn)
	res, err := client.Update(context.Background(), &pb.BookModel{Id: 1, StoreId: 1, Name: "lorem"})
	suite.NotNil(err)
	suite.Nil(res)
	suite.Contains(err.Error(), "lorem")
}

// DESTROY
func (suite *bookGRPCHandlerTestSuite) TestHandler_Destroy_ShouldSuccess() {
	conn := GetConnWithMock(suite, MockDestroySuccess)
	client := pb.NewBookGRPCHandlerClient(conn)
	res, err := client.Destroy(context.Background(), &pb.BookIDModel{Type: pb.ActionType_RELATED, Id: 1})
	suite.Nil(err)
	suite.NotNil(res)
	suite.Equal(res.Status, true)
}
func (suite *bookGRPCHandlerTestSuite) TestHandler_Destroy_ShouldError() {
	conn := GetConnWithMock(suite, MockDestroyError)
	client := pb.NewBookGRPCHandlerClient(conn)
	res, err := client.Destroy(context.Background(), &pb.BookIDModel{Type: pb.ActionType_SPECIFIED, Id: 1})
	suite.NotNil(err)
	suite.Nil(res)
	suite.Contains(err.Error(), "lorem")
}

func TestBookGRPCHandlerService(t *testing.T) {
	suite.Run(t, new(bookGRPCHandlerTestSuite))
}
