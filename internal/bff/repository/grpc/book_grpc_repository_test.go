package grpc_test

import (
	"context"
	"github.com/aasumitro/goms/internal/bff/domain/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type bookGRPCRepositoryTestSuite struct {
	suite.Suite
}

// ALL
func (suite *bookGRPCRepositoryTestSuite) TestHandler_ALL_ShouldSuccess() {
	mockSuite := mockBookGRPCConnection(suite, MockFetchSuccess, "")
	data, err := mockSuite.All(context.TODO(), &entity.Book{StoreID: 1})
	suite.NotNil(data)
	suite.Nil(err)
}
func (suite *bookGRPCRepositoryTestSuite) TestHandler_ALL_ShouldError() {
	mockSuite := mockBookGRPCConnection(suite, MockFetchError, "")
	data, err := mockSuite.All(context.TODO(), nil)
	suite.Nil(data)
	suite.NotNil(err)
	suite.Contains(err.Error(), "lorem")
}

// SHOW
func (suite *bookGRPCRepositoryTestSuite) TestHandler_SHOW_ShouldSuccess() {
	mockSuite := mockBookGRPCConnection(suite, MockShowSuccess, "")
	data, err := mockSuite.Find(context.TODO(), &entity.Book{ID: 1})
	suite.NotNil(data)
	suite.Nil(err)
}
func (suite *bookGRPCRepositoryTestSuite) TestHandler_SHOW_ShouldError() {
	mockSuite := mockBookGRPCConnection(suite, MockShowError, "")
	data, err := mockSuite.Find(context.TODO(), &entity.Book{StoreID: 1})
	suite.Nil(data)
	suite.NotNil(err)
	suite.Contains(err.Error(), "lorem")
}

// RECORD
func (suite *bookGRPCRepositoryTestSuite) TestHandler_RECORD_ShouldSuccess() {
	mockSuite := mockBookGRPCConnection(suite, MockStoreSuccess, "Store")
	err := mockSuite.Record(context.TODO(), &entity.Book{StoreID: 1, Name: "lorem"})
	suite.Nil(err)
}
func (suite *bookGRPCRepositoryTestSuite) TestHandler_RECORD_ShouldError() {
	mockSuite := mockBookGRPCConnection(suite, MockStoreError, "Store")
	err := mockSuite.Record(context.TODO(), &entity.Book{StoreID: 1, Name: "ipsum"})
	suite.NotNil(err)
	suite.Contains(err.Error(), "lorem")
}

// PATCH
func (suite *bookGRPCRepositoryTestSuite) TestHandler_PATCH_ShouldSuccess() {
	mockSuite := mockBookGRPCConnection(suite, MockUpdateSuccess, "Update")
	err := mockSuite.Patch(context.TODO(), &entity.Book{ID: 1, StoreID: 1, Name: "lorem"})
	suite.Nil(err)
}
func (suite *bookGRPCRepositoryTestSuite) TestHandler_PATCH_ShouldError() {
	mockSuite := mockBookGRPCConnection(suite, MockUpdateError, "Update")
	err := mockSuite.Patch(context.TODO(), &entity.Book{ID: 1, StoreID: 1, Name: "lorem"})
	suite.NotNil(err)
	suite.Contains(err.Error(), "lorem")
}

// DELETE
func (suite *bookGRPCRepositoryTestSuite) TestHandler_ERASE_ShouldSuccess() {
	mockSuite := mockBookGRPCConnection(suite, MockDestroySuccess, "Destroy")
	err := mockSuite.Erase(context.TODO(), &entity.Book{ID: 1})
	suite.Nil(err)
}
func (suite *bookGRPCRepositoryTestSuite) TestHandler_ERASE_ShouldError() {
	mockSuite := mockBookGRPCConnection(suite, MockDestroyError, "Destroy")
	err := mockSuite.Erase(context.TODO(), &entity.Book{StoreID: 1})
	suite.NotNil(err)
	suite.Contains(err.Error(), "lorem")
}

func TestBookRepository(t *testing.T) {
	suite.Run(t, new(bookGRPCRepositoryTestSuite))
}
