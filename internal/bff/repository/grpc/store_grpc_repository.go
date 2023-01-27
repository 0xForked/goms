package grpc

import (
	"context"
	"github.com/aasumitro/goms/internal/bff/domain/contract"
	"github.com/aasumitro/goms/internal/bff/domain/entity"
	"github.com/aasumitro/goms/pkg/pb"
)

type storeGRPCRepository struct {
	grpcConn pb.StoreGRPCHandlerClient
}

func (r storeGRPCRepository) All(ctx context.Context) (items []*entity.Store, err error) {
	data, err := r.grpcConn.Fetch(ctx, &pb.StoreEmptyRequest{})
	if err != nil {
		return nil, err
	}

	for _, item := range data.Stores {
		items = append(items, &entity.Store{
			ID:   item.Id,
			Name: item.Name,
		})
	}

	return items, err
}

func (r storeGRPCRepository) Find(ctx context.Context, arg *entity.Store) (item *entity.Store, err error) {
	data, err := r.grpcConn.Show(ctx, &pb.StoreIDModel{Id: arg.ID})
	if err != nil {
		return nil, err
	}

	return &entity.Store{
		ID:   data.Store.Id,
		Name: data.Store.Name,
	}, nil
}

func (r storeGRPCRepository) Record(ctx context.Context, arg *entity.Store) error {
	_, err := r.grpcConn.Store(ctx, &pb.StoreNameModel{Name: arg.Name})
	return err
}

func (r storeGRPCRepository) Patch(ctx context.Context, arg *entity.Store) error {
	_, err := r.grpcConn.Update(ctx, &pb.StoreModel{Id: arg.ID, Name: arg.Name})
	return err
}

func (r storeGRPCRepository) Erase(ctx context.Context, arg *entity.Store) error {
	_, err := r.grpcConn.Destroy(ctx, &pb.StoreIDModel{Id: arg.ID})
	return err
}

func NewStoreGRPCRepository(grpcConn pb.StoreGRPCHandlerClient) contract.IStoreGRPCRepository {
	return &storeGRPCRepository{grpcConn: grpcConn}
}
