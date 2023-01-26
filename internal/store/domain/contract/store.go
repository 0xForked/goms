package contract

import (
	"context"
	"github.com/aasumitro/goms/internal/store/domain/entity"
)

type IStoreRepository interface {
	Select(ctx context.Context, args ...string) (items []*entity.Store, err error)
	Insert(ctx context.Context, args ...*entity.Store) error
	Update(ctx context.Context, arg *entity.Store) error
	Delete(ctx context.Context, arg *entity.Store) error
}

type IStoreService interface {
	All(ctx context.Context) (items []*entity.Store, err error)
	Find(ctx context.Context, arg *entity.Store) (item *entity.Store, err error)
	Record(ctx context.Context, args ...*entity.Store) error
	Patch(ctx context.Context, arg *entity.Store) error
	Erase(ctx context.Context, arg *entity.Store) error
}
