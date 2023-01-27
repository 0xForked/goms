package bff

import (
	"github.com/aasumitro/goms/internal/bff/delivery/handler/rest"
	grpcRepo "github.com/aasumitro/goms/internal/bff/repository/grpc"
	"github.com/aasumitro/goms/internal/bff/service"
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
	storeRepo := grpcRepo.NewStoreGRPCRepository(storeConn)
	bookRepo := grpcRepo.NewBookGRPCRepository(bookConn)
	bffService := service.NewBFFService(redisConn, storeRepo, bookRepo)
	v1 := router.Group("/api/v1")
	rest.NewBFFRegistrarHandler(bffService, v1)
}
