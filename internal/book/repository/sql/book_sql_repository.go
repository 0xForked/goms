package sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aasumitro/goms/internal/book/domain/contract"
	"github.com/aasumitro/goms/internal/book/domain/entity"
)

type bookSQLRepository struct {
	db *sql.DB
}

func (r bookSQLRepository) Select(ctx context.Context, args ...string) (items []*entity.Book, err error) {
	query := "SELECT * FROM books"
	if len(args) > 0 {
		for _, arg := range args {
			query += fmt.Sprintf(" %s", arg)
		}
	}

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	for rows.Next() {
		var book entity.Book
		if err := rows.Scan(&book.ID, &book.StoreID, &book.Name); err != nil {
			return nil, err
		}
		items = append(items, &book)
	}

	return items, nil
}

func (r bookSQLRepository) Insert(ctx context.Context, args ...*entity.Book) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO books (store_id, name) VALUES"
	if len(args) > 0 {
		for i, arg := range args {
			query += fmt.Sprintf(" (%d, '%s')", arg.StoreID, arg.Name)
			if i != (len(args) - 1) {
				query += ","
			}
		}
	}

	if _, err := tx.ExecContext(ctx, query); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r bookSQLRepository) Update(ctx context.Context, arg *entity.Book) error {
	var book entity.Book
	query := "UPDATE books SET name = ?, store_id = ? WHERE id = ? RETURNING *"

	if err := r.db.QueryRowContext(
		ctx, query, arg.Name, arg.StoreID, arg.ID,
	).Scan(
		&book.ID, &book.StoreID, &book.Name,
	); err != nil {
		return err
	}

	return nil
}

func (r bookSQLRepository) Delete(ctx context.Context, arg *entity.Book) error {
	getQuery := func() string {
		query := "DELETE FROM books WHERE "

		if arg.ID != 0 {
			query += "id = ?"
			return query
		}

		query += "store_id = ?"
		return query
	}()

	getID := func() uint32 {
		if arg.ID != 0 {
			return arg.ID
		}
		return arg.StoreID
	}()

	_, err := r.db.ExecContext(ctx, getQuery, getID)
	return err
}

func NewBookSQLRepository(db *sql.DB) contract.IBookRepository {
	return bookSQLRepository{db}
}
