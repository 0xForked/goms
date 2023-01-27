package sql_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/goms/internal/book/domain/contract"
	"github.com/aasumitro/goms/internal/book/domain/entity"
	sqlRepo "github.com/aasumitro/goms/internal/book/repository/sql"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

var (
	err error
	db  *sql.DB
)

type bookSQLRepositoryTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo contract.IBookRepository
}

func (suite *bookSQLRepositoryTestSuite) SetupSuite() {
	db, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))

	require.NoError(suite.T(), err)

	suite.repo = sqlRepo.NewBookSQLRepository(db)
}

// SELECT TEST CASE
func (suite *bookSQLRepositoryTestSuite) TestRepository_Select_ByID_ExpectedReturnRows() {
	mockData := suite.mock.
		NewRows([]string{"id", "store_id", "name"}).
		AddRow(1, 1, "test1").
		AddRow(2, 2, "test2")
	query := "SELECT * FROM books"
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(mockData)
	res, err := suite.repo.Select(context.TODO())
	suite.Nil(err)
	suite.NotNil(res)
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *bookSQLRepositoryTestSuite) TestRepository_Select_ByRelationID_ExpectedReturnRows() {
	mockData := suite.mock.
		NewRows([]string{"id", "store_id", "name"}).
		AddRow(1, 1, "test1").
		AddRow(2, 2, "test2")
	query := "SELECT * FROM books WHERE store_id = 2"
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(mockData)
	res, err := suite.repo.Select(context.TODO(), "WHERE store_id = 2")
	suite.Nil(err)
	suite.NotNil(res)
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *bookSQLRepositoryTestSuite) TestRepository_Select_ExpectedReturnRow() {
	mockData := suite.mock.
		NewRows([]string{"id", "store_id", "name"}).
		AddRow(1, 1, "test1")
	query := "SELECT * FROM books WHERE id = 1"
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(mockData)
	res, err := suite.repo.Select(context.TODO(), "WHERE id = 1")
	suite.Nil(err)
	suite.NotNil(res)
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *bookSQLRepositoryTestSuite) TestRepository_Select_ExpectedErrorFromQuery() {
	query := "SELECT * FROM books"
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).
		WillReturnError(errors.New(""))
	res, err := suite.repo.Select(context.TODO())
	suite.Nil(res)
	suite.NotNil(err)
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *bookSQLRepositoryTestSuite) TestRepository_Select_ExpectedErrorFromScan() {
	mockData := suite.mock.
		NewRows([]string{"id", "store_id", "name"}).
		AddRow(nil, nil, nil)
	query := "SELECT * FROM books"
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(mockData)
	res, err := suite.repo.Select(context.TODO())
	suite.Nil(res)
	suite.NotNil(err)
	suite.Nil(suite.mock.ExpectationsWereMet())
}

// INSERT TEST CASE
func (suite *bookSQLRepositoryTestSuite) TestRepository_Insert_ExpectedSuccess() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(`INSERT INTO books`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()
	err := suite.repo.Insert(context.TODO(), &entity.Book{StoreID: 1, Name: "ipsum"})
	suite.Nil(err)
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *bookSQLRepositoryTestSuite) TestRepository_InsertMultiple_ExpectedSuccess() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(`INSERT INTO books`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()
	err := suite.repo.Insert(context.TODO(),
		&entity.Book{Name: "ipsum", StoreID: 1},
		&entity.Book{Name: "lorem", StoreID: 2})
	suite.Nil(err)
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *bookSQLRepositoryTestSuite) TestRepository_Create_ShouldErrorTxBegin() {
	suite.mock.ExpectBegin().WillReturnError(errors.New("UNEXPECTED"))
	err := suite.repo.Insert(context.TODO(), &entity.Book{Name: "ipsum", StoreID: 1})
	suite.NotNil(err)
	suite.Equal(err.Error(), "UNEXPECTED")
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *bookSQLRepositoryTestSuite) TestRepository_Create_ShouldErrorTxExec() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(`INSERT INTO books`).
		WillReturnError(errors.New("UNEXPECTED"))
	suite.mock.ExpectRollback()
	err := suite.repo.Insert(context.TODO(), &entity.Book{Name: "ipsum", StoreID: 1})
	suite.NotNil(err)
	suite.Equal(err.Error(), "UNEXPECTED")
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *bookSQLRepositoryTestSuite) TestRepository_Create_ShouldErrorCommit() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(`INSERT INTO books`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit().WillReturnError(errors.New("UNEXPECTED"))
	err := suite.repo.Insert(context.TODO(), &entity.Book{Name: "ipsum", StoreID: 1})
	suite.NotNil(err)
	suite.Equal(err.Error(), "UNEXPECTED")
	suite.Nil(suite.mock.ExpectationsWereMet())
}

// UPDATE TEST CASE
func (suite *bookSQLRepositoryTestSuite) TestRepository_Updated_ExpectSuccess() {
	book := &entity.Book{ID: 1, StoreID: 1, Name: "test"}
	data := suite.mock.
		NewRows([]string{"id", "store_id", "name"}).
		AddRow(1, 1, "test")
	query := "UPDATE books SET name = ?, store_id = ? WHERE id = ? RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(book.Name, book.StoreID, book.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	err := suite.repo.Update(context.TODO(), book)
	suite.Nil(err)
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *bookSQLRepositoryTestSuite) TestRepository_Updated_ExpectError() {
	book := &entity.Book{ID: 1, StoreID: 1, Name: "test"}
	data := suite.mock.
		NewRows([]string{"id", "store_id", "name"}).
		AddRow(nil, nil, nil)
	query := "UPDATE books SET name = ?, store_id = ? WHERE id = ? RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(book.Name, book.StoreID, book.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	err := suite.repo.Update(context.TODO(), book)
	suite.NotNil(err)
	suite.Nil(suite.mock.ExpectationsWereMet())
}

// DELETE TEST CASE
func (suite *bookSQLRepositoryTestSuite) TestRepository_Delete_ExpectSuccess() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM books WHERE id = ?")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	err := suite.repo.Delete(context.TODO(), &entity.Book{ID: 1})
	suite.Nil(err)
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *bookSQLRepositoryTestSuite) TestRepository_Delete_ExpectError() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM books WHERE store_id = ?")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnError(errors.New("TEST"))
	err := suite.repo.Delete(context.TODO(), &entity.Book{StoreID: 1})
	suite.NotNil(err)
	suite.Equal(err.Error(), "TEST")
	suite.Nil(suite.mock.ExpectationsWereMet())
}

func TestBookRepository(t *testing.T) {
	suite.Run(t, new(bookSQLRepositoryTestSuite))
}
