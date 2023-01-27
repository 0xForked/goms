package service

import (
	"context"
	"github.com/aasumitro/goms/internal/bff/domain/contract"
	"github.com/aasumitro/goms/internal/bff/domain/entity"
	"github.com/aasumitro/goms/internal/bff/utils"
	"github.com/go-redis/redis/v8"
)

type bffService struct {
	redisConn *redis.Client
	storeRepo contract.IStoreGRPCRepository
	bookRepo  contract.IBookGRPCRepository
}

func (s bffService) AllStore(
	ctx context.Context,
	args *contract.WithParam,
	params *entity.Store,
) (
	items []*entity.Store,
	errorData *utils.ServiceErrorData,
) {
	return utils.WrapDataRows(s.storeRepo.All(ctx))
}

func (s bffService) FirstStore(
	ctx context.Context,
	args *contract.WithParam,
	param *entity.Store,
) (
	items *entity.Store,
	errorData *utils.ServiceErrorData,
) {
	data, err := s.storeRepo.Find(ctx, param)

	if args != nil && *args == contract.WithRelationID {
		if books, err := s.bookRepo.All(ctx, &entity.Book{StoreID: data.ID}); err == nil {
			data.Books = books
		}
	}

	return utils.WrapDataRow(data, err)
}

func NewBFFService(
	redisClient *redis.Client,
	storeRepo contract.IStoreGRPCRepository,
	bookRepo contract.IBookGRPCRepository,
) contract.IBFFService {
	return &bffService{
		redisConn: redisClient,
		storeRepo: storeRepo,
		bookRepo:  bookRepo,
	}
}

// if err := redis.Publish(ctx,
// "notify",
// fmt.Sprintf("send message: %d", i)).
// Err(); err != nil {
// panic(err)
// }

//
// if rows, err := storeRepo.All(context.TODO()); err == nil {
//	for _, row := range rows {
//		fmt.Println("FETCH", row)
//	}
// }

// if row, err := storeRepo.Find(context.TODO(), &entity.Store{ID: 1}); err == nil {
//	fmt.Println("ROW", row)
// }

// if err := storeRepo.Record(context.TODO(), &entity.Store{Name: "lorem"}); err != nil {
//	fmt.Println(err.Error())
// }

// if err := storeRepo.Patch(context.TODO(), &entity.Store{ID: 4, Name: "ipsum"}); err != nil {
//	fmt.Println(err.Error())
// }

// if err := storeRepo.Erase(context.TODO(), &entity.Store{ID: 4}); err != nil {
//	fmt.Println(err.Error())
// }

// bookRepo := grpcRepo.NewBookGRPCRepository(bookConn)
// if rows, err := bookRepo.All(context.TODO(), nil); err == nil {
//	for _, row := range rows {
//		fmt.Println("FETCH_SPECIFIED", row)
//	}
// }

// if rowsRelation, err := bookRepo.All(context.TODO(), &entity.Book{StoreID: 2}); err == nil {
//	for _, row := range rowsRelation {
//		fmt.Println("FETCH_RELATION", row)
//	}
// }

// if row, err := bookRepo.Find(context.TODO(), &entity.Book{ID: 1}); err == nil {
//	fmt.Println("ROW", row)
// }

// if err := bookRepo.Record(context.TODO(), &entity.Book{StoreID: 3, Name: "ipsum"}); err != nil {
//	fmt.Println(err.Error())
// }

// if err := bookRepo.Patch(context.TODO(), &entity.Book{ID: 6, StoreID: 1, Name: "ipsum"}); err != nil {
//	fmt.Println(err.Error())
// }

// if err := bookRepo.Erase(context.TODO(), &entity.Book{StoreID: 3}); err != nil {
//	fmt.Println(err.Error())
// }
