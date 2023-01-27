package grpc_test

import (
	"context"
	"github.com/aasumitro/goms/internal/bff/domain/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type storeGRPCRepositoryTestSuite struct {
	suite.Suite
}

// ALL
func (suite *storeGRPCRepositoryTestSuite) TestHandler_ALL_ShouldSuccess() {
	mockSuite := mockStoreGRPCConnection(suite, MockFetchSuccess, "")
	data, err := mockSuite.All(context.TODO())
	suite.NotNil(data)
	suite.Nil(err)
}
func (suite *storeGRPCRepositoryTestSuite) TestHandler_ALL_ShouldError() {
	mockSuite := mockStoreGRPCConnection(suite, MockFetchError, "")
	data, err := mockSuite.All(context.TODO())
	suite.Nil(data)
	suite.NotNil(err)
	suite.Contains(err.Error(), "lorem")
}

// SHOW
func (suite *storeGRPCRepositoryTestSuite) TestHandler_SHOW_ShouldSuccess() {
	mockSuite := mockStoreGRPCConnection(suite, MockShowSuccess, "")
	data, err := mockSuite.Find(context.TODO(), &entity.Store{ID: 1})
	suite.NotNil(data)
	suite.Nil(err)
}
func (suite *storeGRPCRepositoryTestSuite) TestHandler_SHOW_ShouldError() {
	mockSuite := mockStoreGRPCConnection(suite, MockShowError, "")
	data, err := mockSuite.Find(context.TODO(), &entity.Store{ID: 1})
	suite.Nil(data)
	suite.NotNil(err)
	suite.Contains(err.Error(), "lorem")
}

// RECORD
func (suite *storeGRPCRepositoryTestSuite) TestHandler_RECORD_ShouldSuccess() {
	mockSuite := mockStoreGRPCConnection(suite, MockStoreSuccess, "Store")
	err := mockSuite.Record(context.TODO(), &entity.Store{Name: "lorem"})
	suite.Nil(err)
}
func (suite *storeGRPCRepositoryTestSuite) TestHandler_RECORD_ShouldError() {
	mockSuite := mockStoreGRPCConnection(suite, MockStoreError, "Store")
	err := mockSuite.Record(context.TODO(), &entity.Store{Name: "ipsum"})
	suite.NotNil(err)
	suite.Contains(err.Error(), "lorem")
}

// PATCH
func (suite *storeGRPCRepositoryTestSuite) TestHandler_PATCH_ShouldSuccess() {
	mockSuite := mockStoreGRPCConnection(suite, MockUpdateSuccess, "Update")
	err := mockSuite.Patch(context.TODO(), &entity.Store{ID: 1, Name: "lorem"})
	suite.Nil(err)
}
func (suite *storeGRPCRepositoryTestSuite) TestHandler_PATCH_ShouldError() {
	mockSuite := mockStoreGRPCConnection(suite, MockUpdateError, "Update")
	err := mockSuite.Patch(context.TODO(), &entity.Store{ID: 1, Name: "lorem"})
	suite.NotNil(err)
	suite.Contains(err.Error(), "lorem")
}

// DELETE
func (suite *storeGRPCRepositoryTestSuite) TestHandler_ERASE_ShouldSuccess() {
	mockSuite := mockStoreGRPCConnection(suite, MockDestroySuccess, "Destroy")
	err := mockSuite.Erase(context.TODO(), &entity.Store{ID: 1})
	suite.Nil(err)
}
func (suite *storeGRPCRepositoryTestSuite) TestHandler_ERASE_ShouldError() {
	mockSuite := mockStoreGRPCConnection(suite, MockDestroyError, "Destroy")
	err := mockSuite.Erase(context.TODO(), &entity.Store{ID: 1})
	suite.NotNil(err)
	suite.Contains(err.Error(), "lorem")
}

func TestStoreRepository(t *testing.T) {
	suite.Run(t, new(storeGRPCRepositoryTestSuite))
}
