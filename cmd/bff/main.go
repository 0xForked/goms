package main

import (
	"context"
	"fmt"
	"github.com/aasumitro/goms/pkg/pb"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
)

const RedisAddress = "localhost:6379"

var (
	redisOnce sync.Once
	redisPool *redis.Client
)

func init() {
	getRedisConn()
}

func main() {
	//storeService()
	//bookService()
	publishNotify()
}

func getRedisConn() {
	redisOnce.Do(func() {
		redisPool = redis.NewClient(&redis.Options{Addr: RedisAddress})
		if err := redisPool.Ping(context.Background()).Err(); err != nil {
			panic(fmt.Sprintf(
				"REDIS_ERROR: %s",
				err.Error()))
		}
	})
}

func storeService() {
	var (
		opts []grpc.DialOption
		conn *grpc.ClientConn
		err  error
	)

	opts = append(opts, grpc.WithTransportCredentials(
		insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())
	if conn, err = grpc.Dial(":8002", opts...); err != nil {
		log.Fatalln(fmt.Sprintf("error in dial: %s", err.Error()))
	}
	defer func() { _ = conn.Close() }()

	client := pb.NewStoreGRPCHandlerClient(conn)
	data, err := client.Fetch(context.TODO(), &pb.StoreEmptyRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}

func bookService() {
	var (
		opts []grpc.DialOption
		conn *grpc.ClientConn
		err  error
	)

	opts = append(opts, grpc.WithTransportCredentials(
		insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())
	if conn, err = grpc.Dial(":8002", opts...); err != nil {
		log.Fatalln(fmt.Sprintf("error in dial: %s", err.Error()))
	}
	defer func() { _ = conn.Close() }()

	client := pb.NewBookGRPCHandlerClient(conn)
	data, err := client.Fetch(context.TODO(), &pb.BookIDModel{Type: pb.ActionType_RELATED, Id: 2})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}

func publishNotify() {
	ctx := context.Background()
	for i := 0; i <= 10; i++ {
		if err := redisPool.Publish(ctx, "notify", fmt.Sprintf("send message: %d", i)).Err(); err != nil {
			panic(err)
		}
	}
}
