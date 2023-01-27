package store

import (
	"database/sql"
	handler "github.com/aasumitro/goms/internal/store/delivery/handler/grpc"
	sqlRepo "github.com/aasumitro/goms/internal/store/repository/sql"
	"github.com/aasumitro/goms/internal/store/service"
	"github.com/aasumitro/goms/pkg/pb"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcOpentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewStoreService(db *sql.DB, host net.Listener) {
	repo := sqlRepo.NewStoreSQLRepository(db)
	svc := service.NewStoreService(repo)
	tpt := handler.NewStoreGRPCHandler(svc)
	svr := grpc.NewServer(
		grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			grpcCtxTags.StreamServerInterceptor(),
			grpcOpentracing.StreamServerInterceptor(),
			grpcRecovery.StreamServerInterceptor(),
			// grpcZap.StreamServerInterceptor(zapLogger, []grpcZap.Option{
			//	grpcZap.WithDurationField(func(duration time.Duration) zapcore.Field {
			//		return zap.Int64("grpc.time_ns", duration.Nanoseconds())
			//	}),
			// }...),
		)),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcCtxTags.UnaryServerInterceptor(),
			grpcOpentracing.UnaryServerInterceptor(),
			grpcRecovery.UnaryServerInterceptor(),
			// grpcZap.UnaryServerInterceptor(zapLogger, []grpcZap.Option{
			//	grpcZap.WithDurationField(func(duration time.Duration) zapcore.Field {
			//		return zap.Int64("grpc.time_ns", duration.Nanoseconds())
			//	}),
			// }...),
		)),
	)
	pb.RegisterStoreGRPCHandlerServer(svr, tpt)
	if err := svr.Serve(host); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
