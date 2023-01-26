package main_test

import (
	"context"
	"fmt"
	"github.com/aasumitro/goms/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

func GetGRPCConn(port string) *grpc.ClientConn {
	var (
		opts []grpc.DialOption
		conn *grpc.ClientConn
		err  error
	)

	opts = append(opts, grpc.WithTransportCredentials(
		insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())
	if conn, err = grpc.Dial(port, opts...); err != nil {
		log.Fatalln(fmt.Sprintf("error in dial: %s", err.Error()))
	}
	defer func() { _ = conn.Close() }()

	return conn
}

func Benchmark_StoreServer(b *testing.B) {
	conn := GetGRPCConn(":8001")
	client := pb.NewStoreGRPCHandlerClient(conn)
	for i := 0; i < b.N; i++ {
		_, _ = client.Fetch(context.TODO(), &pb.StoreEmptyRequest{})
	}
}

func Benchmark_BookServer(b *testing.B) {
	conn := GetGRPCConn(":8002")
	client := pb.NewBookGRPCHandlerClient(conn)
	for i := 0; i < b.N; i++ {
		_, _ = client.Fetch(context.TODO(), &pb.BookIDModel{})
	}
}
