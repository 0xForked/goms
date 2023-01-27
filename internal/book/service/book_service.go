package service

import (
	"context"
	"fmt"
	"github.com/aasumitro/goms/internal/book/domain/contract"
	"github.com/aasumitro/goms/internal/book/domain/entity"
)

type bookService struct {
	repo contract.IBookRepository
}

func (s bookService) All(ctx context.Context, args ...string) (items []*entity.Book, err error) {
	return s.repo.Select(ctx, args...)
}

func (s bookService) Find(ctx context.Context, arg *entity.Book) (item *entity.Book, err error) {
	var args string
	if arg.ID != 0 {
		args = fmt.Sprintf(" WHERE id = %d", arg.ID)
	}

	data, err := s.repo.Select(ctx, args)
	if err != nil {
		return nil, err
	}

	return func() *entity.Book {
		if len(data) == 0 {
			return nil
		}
		return data[0]
	}(), nil
}

func (s bookService) Record(ctx context.Context, args ...*entity.Book) error {
	return s.repo.Insert(ctx, args...)
}

func (s bookService) Patch(ctx context.Context, arg *entity.Book) error {
	return s.repo.Update(ctx, arg)
}

func (s bookService) Erase(ctx context.Context, arg *entity.Book) error {
	return s.repo.Delete(ctx, arg)
}

func NewBookService(
	repo contract.IBookRepository,
) contract.IBookService {
	return &bookService{
		repo: repo,
	}
}
