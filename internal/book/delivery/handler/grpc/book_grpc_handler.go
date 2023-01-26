package grpc

import (
	"context"
	"fmt"
	"github.com/aasumitro/goms/internal/book/domain/contract"
	"github.com/aasumitro/goms/internal/book/domain/entity"
	"github.com/aasumitro/goms/pkg/pb"
	"sync"
)

type BookGRPCHandler struct {
	pb.UnimplementedBookGRPCHandlerServer
	mu  sync.Mutex
	Svc contract.IBookService
}

func (handler *BookGRPCHandler) Fetch(
	ctx context.Context,
	model *pb.BookIDModel,
) (*pb.BookRowsResponse, error) {
	var (
		stores []*entity.Book
		err    error
	)

	var args []string
	if model.Type == pb.ActionType_RELATED {
		args = append(args, fmt.Sprintf("WHERE store_id = %d", model.Id))
	}
	if stores, err = handler.Svc.All(ctx, args...); err != nil {
		return nil, err
	}

	var data []*pb.BookModel
	for _, store := range stores {
		data = append(data, &pb.BookModel{
			Id:      store.ID,
			StoreId: store.StoreID,
			Name:    store.Name,
		})
	}

	return &pb.BookRowsResponse{
		Stores: data,
	}, err
}

func (handler *BookGRPCHandler) Show(
	ctx context.Context,
	model *pb.BookIDModel,
) (*pb.BookRowResponse, error) {
	var (
		store *entity.Book
		err   error
	)

	if store, err = handler.Svc.Find(ctx, &entity.Book{ID: model.Id}); err != nil {
		return nil, err
	}

	return &pb.BookRowResponse{Store: &pb.BookModel{
		Id:      store.ID,
		StoreId: store.StoreID,
		Name:    store.Name,
	}}, nil
}

func (handler *BookGRPCHandler) Store(
	ctx context.Context,
	request *pb.BookAddRequest,
) (*pb.BookBoolResponse, error) {
	if err := handler.Svc.Record(ctx, &entity.Book{
		StoreID: request.StoreId,
		Name:    request.Name,
	}); err != nil {
		return &pb.BookBoolResponse{Status: false}, err
	}

	return &pb.BookBoolResponse{Status: true}, nil
}

func (handler *BookGRPCHandler) Update(
	ctx context.Context,
	model *pb.BookModel,
) (*pb.BookBoolResponse, error) {
	if err := handler.Svc.Patch(ctx, &entity.Book{
		ID:      model.Id,
		StoreID: model.StoreId,
		Name:    model.Name,
	}); err != nil {
		return &pb.BookBoolResponse{Status: false}, err
	}

	return &pb.BookBoolResponse{Status: true}, nil
}

func (handler *BookGRPCHandler) Destroy(
	ctx context.Context,
	model *pb.BookIDModel,
) (*pb.BookBoolResponse, error) {
	var book entity.Book
	switch model.Type {
	case pb.ActionType_RELATED:
		book = entity.Book{StoreID: model.Id}
	case pb.ActionType_SPECIFIED:
		book = entity.Book{ID: model.Id}
	}

	if err := handler.Svc.Erase(ctx, &book); err != nil {
		return &pb.BookBoolResponse{Status: false}, err
	}

	return &pb.BookBoolResponse{Status: true}, nil
}

func NewBookGRPCHandler(service contract.IBookService) *BookGRPCHandler {
	return &BookGRPCHandler{Svc: service}
}
