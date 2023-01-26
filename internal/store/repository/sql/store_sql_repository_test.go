package sql_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bakode/goms/internal/store/domain/contract"
	"github.com/bakode/goms/internal/store/domain/entity"
	sqlRepo "github.com/bakode/goms/internal/store/repository/sql"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

var (
	err error
	db  *sql.DB
)

type storeSQLRepositoryTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo contract.IStoreRepository
}

func (suite *storeSQLRepositoryTestSuite) SetupSuite() {
	db, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))

	require.NoError(suite.T(), err)

	suite.repo = sqlRepo.NewStoreSQLRepository(db)
}

// SELECT TEST CASE
func (suite *storeSQLRepositoryTestSuite) TestRepository_Select_ExpectedReturnRows() {
	mockData := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(1, "test1").
		AddRow(2, "test2")
	query := "SELECT * FROM stores"
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(mockData)
	res, err := suite.repo.Select(context.TODO())
	suite.Nil(err)
	suite.NotNil(res)
	suite.Nil(suite.mock.ExpectationsWereMet())

}
func (suite *storeSQLRepositoryTestSuite) TestRepository_Select_ExpectedReturnRow() {
	mockData := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(1, "test1")
	query := "SELECT * FROM stores WHERE id = 1"
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(mockData)
	res, err := suite.repo.Select(context.TODO(), "WHERE id = 1")
	suite.Nil(err)
	suite.NotNil(res)
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *storeSQLRepositoryTestSuite) TestRepository_Select_ExpectedErrorFromQuery() {
	query := "SELECT * FROM stores"
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).
		WillReturnError(errors.New(""))
	res, err := suite.repo.Select(context.TODO())
	suite.Nil(res)
	suite.NotNil(err)
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *storeSQLRepositoryTestSuite) TestRepository_Select_ExpectedErrorFromScan() {
	mockData := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(nil, nil)
	query := "SELECT * FROM stores"
	expectedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(mockData)
	res, err := suite.repo.Select(context.TODO())
	suite.Nil(res)
	suite.NotNil(err)
	suite.Nil(suite.mock.ExpectationsWereMet())
}

// INSERT TEST CASE
func (suite *storeSQLRepositoryTestSuite) TestRepository_Insert_ExpectedSuccess() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(`INSERT INTO stores`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()
	err := suite.repo.Insert(context.TODO(), &entity.Store{Name: "ipsum"})
	suite.Nil(err)
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *storeSQLRepositoryTestSuite) TestRepository_InsertMultiple_ExpectedSuccess() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(`INSERT INTO stores`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()
	err := suite.repo.Insert(context.TODO(), &entity.Store{Name: "ipsum"}, &entity.Store{Name: "lorem"})
	suite.Nil(err)
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *storeSQLRepositoryTestSuite) TestRepository_Create_ShouldErrorTxBegin() {
	suite.mock.ExpectBegin().WillReturnError(errors.New("UNEXPECTED"))
	err := suite.repo.Insert(context.TODO(), &entity.Store{Name: "ipsum"})
	suite.NotNil(err)
	suite.Equal(err.Error(), "UNEXPECTED")
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *storeSQLRepositoryTestSuite) TestRepository_Create_ShouldErrorTxExec() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(`INSERT INTO stores`).
		WillReturnError(errors.New("UNEXPECTED"))
	suite.mock.ExpectRollback()
	err := suite.repo.Insert(context.TODO(), &entity.Store{Name: "ipsum"})
	suite.NotNil(err)
	suite.Equal(err.Error(), "UNEXPECTED")
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *storeSQLRepositoryTestSuite) TestRepository_Create_ShouldErrorCommit() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(`INSERT INTO stores`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit().WillReturnError(errors.New("UNEXPECTED"))
	err := suite.repo.Insert(context.TODO(), &entity.Store{Name: "ipsum"})
	suite.NotNil(err)
	suite.Equal(err.Error(), "UNEXPECTED")
	suite.Nil(suite.mock.ExpectationsWereMet())
}

// UPDATE TEST CASE
func (suite *storeSQLRepositoryTestSuite) TestRepository_Updated_ExpectSuccess() {
	store := &entity.Store{ID: 1, Name: "test"}
	data := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(1, "test")
	query := "UPDATE stores SET name = ? WHERE id = ? RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(store.Name, store.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	err := suite.repo.Update(context.TODO(), store)
	suite.Nil(err)
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *storeSQLRepositoryTestSuite) TestRepository_Updated_ExpectError() {
	store := &entity.Store{ID: 1, Name: "test"}
	data := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(nil, nil)
	query := "UPDATE stores SET name = ? WHERE id = ? RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(store.Name, store.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	err := suite.repo.Update(context.TODO(), store)
	suite.NotNil(err)
	suite.Nil(suite.mock.ExpectationsWereMet())
}

// DELETE TEST CASE
func (suite *storeSQLRepositoryTestSuite) TestRepository_Delete_ExpectSuccess() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM stores WHERE id = ?")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	err := suite.repo.Delete(context.TODO(), &entity.Store{ID: 1})
	suite.Nil(err)
	suite.Nil(suite.mock.ExpectationsWereMet())
}
func (suite *storeSQLRepositoryTestSuite) TestRepository_Delete_ExpectError() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM stores WHERE id = ?")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnError(errors.New("TEST"))
	err := suite.repo.Delete(context.TODO(), &entity.Store{ID: 1})
	suite.NotNil(err)
	suite.Equal(err.Error(), "TEST")
	suite.Nil(suite.mock.ExpectationsWereMet())
}

func TestStoreRepository(t *testing.T) {
	suite.Run(t, new(storeSQLRepositoryTestSuite))
}
