package service

import (
	"github.com/aasumitro/goms/internal/bff/domain/contract"
)

type bffService struct {
	storeRepo contract.IStoreGRPCRepository
	bookRepo  contract.IBookGRPCRepository
}

func NewBFFService(
	storeRepo contract.IStoreGRPCRepository,
	bookRepo contract.IBookGRPCRepository,
) contract.IBFFService {
	return &bffService{
		storeRepo: storeRepo,
		bookRepo:  bookRepo,
	}
}

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
