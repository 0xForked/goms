package main

import (
	"context"
	"fmt"
	"github.com/aasumitro/goms/docs"
	"github.com/aasumitro/goms/internal/bff"
	"github.com/aasumitro/goms/pkg/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"sync"
)

const (
	name        = "GOMS"
	description = "Microservice Example with EDA & ADA"
	version     = "v0.0.1"
	production  = false

	serviceAddress      = "localhost:8000"
	storeServiceAddress = "localhost:8001"
	bookServiceAddress  = "localhost:8002"
	redisAddress        = "localhost:6379"

	ginModelsDepth = 4
)

var (
	redisOnce        sync.Once
	storeServiceOnce sync.Once
	bookServiceOnce  sync.Once

	redisPool        *redis.Client
	storeServiceConn pb.StoreGRPCHandlerClient
	bookServiceConn  pb.BookGRPCHandlerClient
	appEngine        *gin.Engine
)

func init() {
	initRedisConn()
	initStoreServiceConn()
	initBookServiceConn()
	initGinEngine()
	initSwaggerInfo()
}

func main() {
	bff.NewBFFService(appEngine, redisPool, storeServiceConn, bookServiceConn)

	log.Fatal(appEngine.Run(serviceAddress))
}

func initRedisConn() {
	redisOnce.Do(func() {
		redisPool = redis.NewClient(&redis.Options{Addr: redisAddress})
		if err := redisPool.Ping(context.Background()).Err(); err != nil {
			panic(fmt.Sprintf(
				"REDIS_ERROR: %s",
				err.Error()))
		}
	})
}

func initStoreServiceConn() {
	storeServiceOnce.Do(func() {
		conn := getGRPCConn(storeServiceAddress)
		storeServiceConn = pb.NewStoreGRPCHandlerClient(conn)
	})
}

func initBookServiceConn() {
	bookServiceOnce.Do(func() {
		conn := getGRPCConn(bookServiceAddress)
		bookServiceConn = pb.NewBookGRPCHandlerClient(conn)
	})
}

func initGinEngine() {
	if production {
		gin.SetMode(gin.ReleaseMode)
	}

	appEngine = gin.Default()

	appEngine.NoMethod(func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/docs/index.html")
	})

	appEngine.NoRoute(func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/docs/index.html")
	})

	appEngine.GET("/docs/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.DefaultModelsExpandDepth(ginModelsDepth)))
}

func initSwaggerInfo() {
	docs.SwaggerInfo.BasePath = appEngine.BasePath()
	docs.SwaggerInfo.Title = name
	docs.SwaggerInfo.Description = description
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = serviceAddress
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func getGRPCConn(add string) *grpc.ClientConn {
	var (
		conn *grpc.ClientConn
		err  error
	)

	if conn, err = grpc.Dial(add,
		grpc.WithTransportCredentials(
			insecure.NewCredentials()),
		grpc.WithBlock(),
	); err != nil {
		panic(fmt.Sprintf(
			"GRPC_DIAL_ERROR: %s",
			err.Error()))
	}

	// defer func() { _ = conn.Close() }()

	return conn
}
