package store

import (
	"database/sql"
	handler "github.com/aasumitro/goms/internal/book/delivery/handler/grpc"
	sqlRepo "github.com/aasumitro/goms/internal/book/repository/sql"
	"github.com/aasumitro/goms/internal/book/service"
	"github.com/aasumitro/goms/pkg/pb"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcOpentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewBookService(db *sql.DB, host net.Listener) {
	repo := sqlRepo.NewBookSQLRepository(db)
	svc := service.NewBookService(repo)
	tpt := handler.NewBookGRPCHandler(svc)
	svr := grpc.NewServer(
		grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			grpcCtxtags.StreamServerInterceptor(),
			grpcOpentracing.StreamServerInterceptor(),
			grpcRecovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcCtxtags.UnaryServerInterceptor(),
			grpcOpentracing.UnaryServerInterceptor(),
			grpcRecovery.UnaryServerInterceptor(),
		)),
	)
	pb.RegisterBookGRPCHandlerServer(svr, tpt)
	if err := svr.Serve(host); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
