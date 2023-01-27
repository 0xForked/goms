package bff

import (
	"context"
	"fmt"
	grpcRepo "github.com/aasumitro/goms/internal/bff/repository/grpc"
	"github.com/aasumitro/goms/pkg/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewBFFService(
	router *gin.Engine,
	redisConn *redis.Client,
	storeConn pb.StoreGRPCHandlerClient,
	bookConn pb.BookGRPCHandlerClient,
) {
	_ = grpcRepo.NewStoreGRPCRepository(storeConn)
	_ = grpcRepo.NewBookGRPCRepository(bookConn)

	publish(redisConn)
}

func publish(redis *redis.Client) {
	ctx := context.Background()
	for i := 0; i <= 10; i++ {
		if err := redis.Publish(ctx,
			"notify",
			fmt.Sprintf("send message: %d", i)).
			Err(); err != nil {
			panic(err)
		}
	}
}
