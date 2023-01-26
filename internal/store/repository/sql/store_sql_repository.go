package sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/bakode/goms/internal/store/domain/contract"
	"github.com/bakode/goms/internal/store/domain/entity"
)

type storeSQLRepository struct {
	db *sql.DB
}

func (r storeSQLRepository) Select(ctx context.Context, args ...string) (items []*entity.Store, err error) {
	query := "SELECT * FROM stores"
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
		var store entity.Store
		if err := rows.Scan(&store.ID, &store.Name); err != nil {
			return nil, err
		}
		items = append(items, &store)
	}

	return items, nil
}

func (r storeSQLRepository) Insert(ctx context.Context, args ...*entity.Store) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO stores (name) VALUES"
	if len(args) > 0 {
		for i, arg := range args {
			query += fmt.Sprintf(" ('%s')", arg.Name)
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

func (r storeSQLRepository) Update(ctx context.Context, arg *entity.Store) error {
	var store entity.Store
	query := "UPDATE stores SET name = ? WHERE id = ? RETURNING *"

	if err := r.db.QueryRowContext(
		ctx, query, arg.Name, arg.ID,
	).Scan(
		&store.ID, &store.Name,
	); err != nil {
		return err
	}

	return nil
}

func (r storeSQLRepository) Delete(ctx context.Context, arg *entity.Store) error {
	query := "DELETE FROM stores WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, arg.ID)
	return err
}

func NewStoreSQLRepository(db *sql.DB) contract.IStoreRepository {
	return storeSQLRepository{db}
}
