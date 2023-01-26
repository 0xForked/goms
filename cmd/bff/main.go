package main

import (
	"context"
	"fmt"
	"github.com/bakode/goms/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
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
