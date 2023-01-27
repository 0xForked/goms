package contract

import (
	"context"
	"github.com/aasumitro/goms/internal/bff/domain/entity"
	"github.com/aasumitro/goms/internal/bff/utils"
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

	WithParam int8

	IBFFService interface {
		AllStore(ctx context.Context, args *WithParam, params *entity.Store) (items []*entity.Store, errorData *utils.ServiceErrorData)
		FirstStore(ctx context.Context, args *WithParam, param *entity.Store) (items *entity.Store, errorData *utils.ServiceErrorData)
		//CreateStore()
		//UpdateStore()
		//DestroyStore()

		//AllBook(ctx context.Context, args *WithParam, params *entity.Book) (items []*entity.Book, errorData *utils.ServiceErrorData)
		//FirstBook(ctx context.Context, param *entity.Book) (items *entity.Book, errorData *utils.ServiceErrorData)
		//CreateBook()
		//UpdateBook()
		//DestroyBook()
	}
)

const (
	WithID WithParam = iota
	WithRelationID
)
