package service_test

import (
	"context"
	"errors"
	"github.com/aasumitro/goms/internal/book/domain/entity"
	"github.com/aasumitro/goms/internal/book/service"
	"github.com/aasumitro/goms/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type bookServiceTestSuite struct {
	suite.Suite
	books []*entity.Book
}

func (suite *bookServiceTestSuite) SetupSuite() {
	suite.books = []*entity.Book{
		{ID: 1, StoreID: 1, Name: "lorem"},
		{ID: 2, StoreID: 2, Name: "ipsum"},
	}
}

// ALL
func (suite *bookServiceTestSuite) TestService_All_ShouldSuccess() {
	repoMock := new(mocks.IBookRepository)
	svc := service.NewBookService(repoMock)
	repoMock.On("Select", mock.Anything).
		Return(suite.books, nil).Once()
	data, err := svc.All(context.TODO())
	suite.Nil(err)
	suite.NotNil(data)
	suite.Equal(data, suite.books)
	repoMock.AssertExpectations(suite.T())
}
func (suite *bookServiceTestSuite) TestService_All_ShouldError() {
	repoMock := new(mocks.IBookRepository)
	svc := service.NewBookService(repoMock)
	repoMock.On("Select", mock.Anything).
		Return(nil, errors.New("lorem")).Once()
	data, err := svc.All(context.TODO())
	suite.NotNil(err)
	suite.Nil(data)
	suite.Equal(err.Error(), "lorem")
	repoMock.AssertExpectations(suite.T())
}

// FIND
func (suite *bookServiceTestSuite) TestService_Find_ShouldSuccess() {
	repoMock := new(mocks.IBookRepository)
	svc := service.NewBookService(repoMock)
	repoMock.On("Select", mock.Anything, mock.Anything).
		Return(suite.books, nil).Once()
	data, err := svc.Find(context.TODO(), &entity.Book{ID: 1})
	suite.Nil(err)
	suite.NotNil(data)
	suite.Equal(data, suite.books[0])
	repoMock.AssertExpectations(suite.T())
}
func (suite *bookServiceTestSuite) TestService_Find_ShouldError() {
	repoMock := new(mocks.IBookRepository)
	svc := service.NewBookService(repoMock)
	repoMock.On("Select", mock.Anything, mock.Anything).
		Return(nil, errors.New("lorem")).Once()
	data, err := svc.Find(context.TODO(), &entity.Book{ID: 1})
	suite.NotNil(err)
	suite.Nil(data)
	suite.Equal(err.Error(), "lorem")
	repoMock.AssertExpectations(suite.T())
}

// CREATE
func (suite *bookServiceTestSuite) TestRepository_Create_ExpectedSuccess() {
	repoMock := new(mocks.IBookRepository)
	svc := service.NewBookService(repoMock)
	repoMock.On("Insert", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := svc.Record(context.TODO(), &entity.Book{ID: 1, StoreID: 1, Name: "lorem"})
	suite.Nil(err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *bookServiceTestSuite) TestRepository_Create_ExpectedError() {
	repoMock := new(mocks.IBookRepository)
	svc := service.NewBookService(repoMock)
	repoMock.On("Insert", mock.Anything, mock.Anything).
		Return(errors.New("lorem")).Once()
	err := svc.Record(context.TODO(), &entity.Book{ID: 1, StoreID: 1, Name: "lorem"})
	suite.NotNil(err)
	suite.Equal(err.Error(), "lorem")
	repoMock.AssertExpectations(suite.T())
}

// UPDATE
func (suite *bookServiceTestSuite) TestRepository_Update_ExpectedSuccess() {
	repoMock := new(mocks.IBookRepository)
	svc := service.NewBookService(repoMock)
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := svc.Patch(context.TODO(), &entity.Book{ID: 1, StoreID: 1, Name: "lorem"})
	suite.Nil(err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *bookServiceTestSuite) TestRepository_Update_ExpectedError() {
	repoMock := new(mocks.IBookRepository)
	svc := service.NewBookService(repoMock)
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(errors.New("lorem")).Once()
	err := svc.Patch(context.TODO(), &entity.Book{ID: 1, StoreID: 1, Name: "lorem"})
	suite.NotNil(err)
	suite.Equal(err.Error(), "lorem")
	repoMock.AssertExpectations(suite.T())
}

// Destroy
func (suite *bookServiceTestSuite) TestRepository_Delete_ExpectedSuccess() {
	repoMock := new(mocks.IBookRepository)
	svc := service.NewBookService(repoMock)
	repoMock.On("Delete", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := svc.Erase(context.TODO(), &entity.Book{ID: 1})
	suite.Nil(err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *bookServiceTestSuite) TestRepository_Delete_ExpectedError() {
	repoMock := new(mocks.IBookRepository)
	svc := service.NewBookService(repoMock)
	repoMock.On("Delete", mock.Anything, mock.Anything).
		Return(errors.New("lorem")).Once()
	err := svc.Erase(context.TODO(), &entity.Book{ID: 1})
	suite.NotNil(err)
	suite.Equal(err.Error(), "lorem")
	repoMock.AssertExpectations(suite.T())
}

func TestBookService(t *testing.T) {
	suite.Run(t, new(bookServiceTestSuite))
}
