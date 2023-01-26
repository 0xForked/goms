package contract

import (
	"context"
	"github.com/bakode/goms/internal/book/domain/entity"
)

type (
	IBookRepository interface {
		Select(ctx context.Context, args ...string) (items []*entity.Book, err error)
		Insert(ctx context.Context, args ...*entity.Book) error
		Update(ctx context.Context, arg *entity.Book) error
		Delete(ctx context.Context, arg *entity.Book) error
	}

	IBookService interface {
		All(ctx context.Context, args ...string) (items []*entity.Book, err error)
		Find(ctx context.Context, arg *entity.Book) (item *entity.Book, err error)
		Record(ctx context.Context, args ...*entity.Book) error
		Patch(ctx context.Context, arg *entity.Book) error
		Erase(ctx context.Context, arg *entity.Book) error
	}
)
