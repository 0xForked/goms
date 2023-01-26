package grpc

import (
	"context"
	"github.com/bakode/goms/internal/store/domain/contract"
	"github.com/bakode/goms/internal/store/domain/entity"
	"github.com/bakode/goms/pkg/pb"
	"sync"
)

type StoreGRPCHandler struct {
	pb.UnimplementedStoreGRPCHandlerServer
	mu  sync.Mutex
	Svc contract.IStoreService
}

func (handler *StoreGRPCHandler) Fetch(
	ctx context.Context,
	_ *pb.StoreEmptyRequest,
) (*pb.StoreRowsResponse, error) {
	var (
		stores []*entity.Store
		err    error
	)

	if stores, err = handler.Svc.All(ctx); err != nil {
		return nil, err
	}

	var data []*pb.StoreModel
	for _, store := range stores {
		data = append(data, &pb.StoreModel{
			Id:   store.ID,
			Name: store.Name,
		})
	}

	return &pb.StoreRowsResponse{
		Stores: data,
	}, err
}

func (handler *StoreGRPCHandler) Show(
	ctx context.Context,
	model *pb.StoreIDModel,
) (*pb.StoreRowResponse, error) {
	var (
		store *entity.Store
		err   error
	)

	if store, err = handler.Svc.Find(ctx, &entity.Store{ID: model.Id}); err != nil {
		return nil, err
	}

	return &pb.StoreRowResponse{Store: &pb.StoreModel{
		Id:   store.ID,
		Name: store.Name,
	}}, nil
}

func (handler *StoreGRPCHandler) Store(
	ctx context.Context,
	model *pb.StoreNameModel,
) (*pb.StoreBoolResponse, error) {
	if err := handler.Svc.Record(ctx, &entity.Store{Name: model.Name}); err != nil {
		return &pb.StoreBoolResponse{Status: false}, err
	}

	return &pb.StoreBoolResponse{Status: true}, nil
}

func (handler *StoreGRPCHandler) Update(
	ctx context.Context,
	model *pb.StoreModel,
) (*pb.StoreBoolResponse, error) {
	if err := handler.Svc.Patch(ctx, &entity.Store{ID: model.Id, Name: model.Name}); err != nil {
		return &pb.StoreBoolResponse{Status: false}, err
	}

	return &pb.StoreBoolResponse{Status: true}, nil
}

func (handler *StoreGRPCHandler) Destroy(
	ctx context.Context,
	model *pb.StoreIDModel,
) (*pb.StoreBoolResponse, error) {
	if err := handler.Svc.Erase(ctx, &entity.Store{ID: model.Id}); err != nil {
		return &pb.StoreBoolResponse{Status: false}, err
	}

	return &pb.StoreBoolResponse{Status: true}, nil
}

func NewStoreGRPCHandler(service contract.IStoreService) *StoreGRPCHandler {
	return &StoreGRPCHandler{Svc: service}
}
