package contract

import (
	"context"
	"github.com/aasumitro/goms/internal/bff/domain/entity"
)

type (
	IStoreGRPCRepository interface {
		All(ctx context.Context) (items []*entity.Store, err error)
		Find(ctx context.Context, arg *entity.Store) (item *entity.Store, err error)
		Record(ctx context.Context, arg *entity.Store) error
		Patch(ctx context.Context, arg *entity.Store) error
		Erase(ctx context.Context, arg *entity.Store) error
	}

	IBookGRPCRepository interface {
		All(ctx context.Context, arg *entity.Book) (items []*entity.Book, err error)
		Find(ctx context.Context, arg *entity.Book) (item *entity.Book, err error)
		Record(ctx context.Context, arg *entity.Book) error
		Patch(ctx context.Context, arg *entity.Book) error
		Erase(ctx context.Context, arg *entity.Book) error
	}

	IBFFService interface {
	}
)
