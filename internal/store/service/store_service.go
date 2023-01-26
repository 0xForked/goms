package service

import (
	"context"
	"fmt"
	"github.com/bakode/goms/internal/store/domain/contract"
	"github.com/bakode/goms/internal/store/domain/entity"
)

type storeService struct {
	repo contract.IStoreRepository
}

func (s storeService) All(ctx context.Context) (items []*entity.Store, err error) {
	return s.repo.Select(ctx)
}

func (s storeService) Find(ctx context.Context, arg *entity.Store) (item *entity.Store, err error) {
	var args string
	if arg.ID != 0 {
		args = fmt.Sprintf(" WHERE id = %d", arg.ID)
	}

	data, err := s.repo.Select(ctx, args)
	if err != nil {
		return nil, err
	}

	return data[0], nil
}

func (s storeService) Record(ctx context.Context, args ...*entity.Store) error {
	return s.repo.Insert(ctx, args...)
}

func (s storeService) Patch(ctx context.Context, arg *entity.Store) error {
	return s.repo.Update(ctx, arg)
}

func (s storeService) Erase(ctx context.Context, arg *entity.Store) error {
	return s.repo.Delete(ctx, arg)
}

func NewStoreService(
	repo contract.IStoreRepository,
) contract.IStoreService {
	return &storeService{
		repo: repo,
	}
}
