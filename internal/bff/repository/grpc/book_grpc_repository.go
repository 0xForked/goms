package grpc

import (
	"context"
	"github.com/aasumitro/goms/internal/bff/domain/contract"
	"github.com/aasumitro/goms/internal/bff/domain/entity"
	"github.com/aasumitro/goms/pkg/pb"
)

type bookGRPCRepository struct {
	grpcConn pb.BookGRPCHandlerClient
}

func (r bookGRPCRepository) All(ctx context.Context, arg *entity.Book) (items []*entity.Book, err error) {
	param := func() *pb.BookIDModel {
		if arg == nil {
			return &pb.BookIDModel{}
		}

		return &pb.BookIDModel{
			Type: pb.ActionType_RELATED,
			Id:   arg.StoreID,
		}
	}()

	data, err := r.grpcConn.Fetch(ctx, param)
	if err != nil {
		return nil, err
	}

	for _, item := range data.Books {
		items = append(items, &entity.Book{
			ID:      item.Id,
			StoreID: item.StoreId,
			Name:    item.Name,
		})
	}

	return items, err
}

func (r bookGRPCRepository) Find(ctx context.Context, arg *entity.Book) (item *entity.Book, err error) {
	data, err := r.grpcConn.Show(ctx, &pb.BookIDModel{
		Id:   arg.ID,
		Type: pb.ActionType_SPECIFIED,
	})
	if err != nil {
		return nil, err
	}

	return &entity.Book{
		ID:      data.Book.Id,
		StoreID: data.Book.StoreId,
		Name:    data.Book.Name,
	}, nil
}

func (r bookGRPCRepository) Record(ctx context.Context, arg *entity.Book) error {
	_, err := r.grpcConn.Store(ctx, &pb.BookAddRequest{
		StoreId: arg.StoreID,
		Name:    arg.Name,
	})

	return err
}

func (r bookGRPCRepository) Patch(ctx context.Context, arg *entity.Book) error {
	_, err := r.grpcConn.Update(ctx, &pb.BookModel{
		Id:      arg.ID,
		StoreId: arg.StoreID,
		Name:    arg.Name,
	})

	return err
}

func (r bookGRPCRepository) Erase(ctx context.Context, arg *entity.Book) error {
	param := func() *pb.BookIDModel {
		if arg.StoreID != 0 {
			return &pb.BookIDModel{
				Type: pb.ActionType_RELATED,
				Id:   arg.StoreID,
			}
		}

		return &pb.BookIDModel{
			Type: pb.ActionType_SPECIFIED,
			Id:   arg.ID,
		}
	}()

	_, err := r.grpcConn.Destroy(ctx, param)

	return err
}

func NewBookGRPCRepository(grpcConn pb.BookGRPCHandlerClient) contract.IBookGRPCRepository {
	return &bookGRPCRepository{grpcConn: grpcConn}
}
