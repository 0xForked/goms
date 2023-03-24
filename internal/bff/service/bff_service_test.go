package service_test

import (
	"context"
	"github.com/aasumitro/goms/internal/bff/service"
	"github.com/aasumitro/goms/mocks"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type bffServiceTestSuite struct {
	suite.Suite
	redisConn *redis.Client
}

func (suite *bffServiceTestSuite) SetupSuite() {
	suite.redisConn = redis.NewClient(&redis.Options{
		Addr: miniredis.RunT(suite.T()).Addr(),
	})
}

func (suite *bffServiceTestSuite) TestService_AllStore() {
	repo := new(mocks.IStoreGRPCRepository)
	svc := service.NewBFFService(suite.redisConn, repo, new(mocks.IBookGRPCRepository))
	repo.On("All", mock.Anything).
		Return(nil, nil).Once()
	data, err := svc.AllStore(context.TODO(), nil, nil)
	suite.Nil(err)
	suite.Nil(data)
	repo.AssertExpectations(suite.T())
}

func TestBFFService(t *testing.T) {
	suite.Run(t, new(bffServiceTestSuite))
}
