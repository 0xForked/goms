package service_test

import (
	"context"
	"errors"
	"github.com/aasumitro/goms/internal/store/domain/entity"
	"github.com/aasumitro/goms/internal/store/service"
	"github.com/aasumitro/goms/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type storeServiceTestSuite struct {
	suite.Suite
	stores []*entity.Store
}

func (suite *storeServiceTestSuite) SetupSuite() {
	suite.stores = []*entity.Store{
		{ID: 1, Name: "lorem"},
		{ID: 2, Name: "ipsum"},
	}
}

// ALL
func (suite *storeServiceTestSuite) TestService_All_ShouldSuccess() {
	repoMock := new(mocks.IStoreRepository)
	svc := service.NewStoreService(repoMock)
	repoMock.On("Select", mock.Anything).
		Return(suite.stores, nil).Once()
	data, err := svc.All(context.TODO())
	suite.Nil(err)
	suite.NotNil(data)
	suite.Equal(data, suite.stores)
	repoMock.AssertExpectations(suite.T())
}
func (suite *storeServiceTestSuite) TestService_All_ShouldError() {
	repoMock := new(mocks.IStoreRepository)
	svc := service.NewStoreService(repoMock)
	repoMock.On("Select", mock.Anything).
		Return(nil, errors.New("lorem")).Once()
	data, err := svc.All(context.TODO())
	suite.NotNil(err)
	suite.Nil(data)
	suite.Equal(err.Error(), "lorem")
	repoMock.AssertExpectations(suite.T())
}

// FIND
func (suite *storeServiceTestSuite) TestService_Find_ShouldSuccess() {
	repoMock := new(mocks.IStoreRepository)
	svc := service.NewStoreService(repoMock)
	repoMock.On("Select", mock.Anything, mock.Anything).
		Return(suite.stores, nil).Once()
	data, err := svc.Find(context.TODO(), &entity.Store{ID: 1})
	suite.Nil(err)
	suite.NotNil(data)
	suite.Equal(data, suite.stores[0])
	repoMock.AssertExpectations(suite.T())
}
func (suite *storeServiceTestSuite) TestService_Find_ShouldError() {
	repoMock := new(mocks.IStoreRepository)
	svc := service.NewStoreService(repoMock)
	repoMock.On("Select", mock.Anything, mock.Anything).
		Return(nil, errors.New("lorem")).Once()
	data, err := svc.Find(context.TODO(), &entity.Store{ID: 1})
	suite.NotNil(err)
	suite.Nil(data)
	suite.Equal(err.Error(), "lorem")
	repoMock.AssertExpectations(suite.T())
}

// CREATE
func (suite *storeServiceTestSuite) TestRepository_Create_ExpectedSuccess() {
	repoMock := new(mocks.IStoreRepository)
	svc := service.NewStoreService(repoMock)
	repoMock.On("Insert", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := svc.Record(context.TODO(), &entity.Store{ID: 1, Name: "lorem"})
	suite.Nil(err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *storeServiceTestSuite) TestRepository_Create_ExpectedError() {
	repoMock := new(mocks.IStoreRepository)
	svc := service.NewStoreService(repoMock)
	repoMock.On("Insert", mock.Anything, mock.Anything).
		Return(errors.New("lorem")).Once()
	err := svc.Record(context.TODO(), &entity.Store{ID: 1, Name: "lorem"})
	suite.NotNil(err)
	suite.Equal(err.Error(), "lorem")
	repoMock.AssertExpectations(suite.T())
}

// UPDATE
func (suite *storeServiceTestSuite) TestRepository_Update_ExpectedSuccess() {
	repoMock := new(mocks.IStoreRepository)
	svc := service.NewStoreService(repoMock)
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := svc.Patch(context.TODO(), &entity.Store{ID: 1, Name: "lorem"})
	suite.Nil(err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *storeServiceTestSuite) TestRepository_Update_ExpectedError() {
	repoMock := new(mocks.IStoreRepository)
	svc := service.NewStoreService(repoMock)
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(errors.New("lorem")).Once()
	err := svc.Patch(context.TODO(), &entity.Store{ID: 1, Name: "lorem"})
	suite.NotNil(err)
	suite.Equal(err.Error(), "lorem")
	repoMock.AssertExpectations(suite.T())
}

// Destroy
func (suite *storeServiceTestSuite) TestRepository_Delete_ExpectedSuccess() {
	repoMock := new(mocks.IStoreRepository)
	svc := service.NewStoreService(repoMock)
	repoMock.On("Delete", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := svc.Erase(context.TODO(), &entity.Store{ID: 1})
	suite.Nil(err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *storeServiceTestSuite) TestRepository_Delete_ExpectedError() {
	repoMock := new(mocks.IStoreRepository)
	svc := service.NewStoreService(repoMock)
	repoMock.On("Delete", mock.Anything, mock.Anything).
		Return(errors.New("lorem")).Once()
	err := svc.Erase(context.TODO(), &entity.Store{ID: 1})
	suite.NotNil(err)
	suite.Equal(err.Error(), "lorem")
	repoMock.AssertExpectations(suite.T())
}

func TestStoreService(t *testing.T) {
	suite.Run(t, new(storeServiceTestSuite))
}
