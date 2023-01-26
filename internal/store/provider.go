package store

import (
	"database/sql"
	handler "github.com/bakode/goms/internal/store/delivery/handler/grpc"
	sqlRepo "github.com/bakode/goms/internal/store/repository/sql"
	"github.com/bakode/goms/internal/store/service"
	"github.com/bakode/goms/pkg/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewStoreService(db *sql.DB, host net.Listener) {
	repo := sqlRepo.NewStoreSQLRepository(db)
	svc := service.NewStoreService(repo)
	tpt := handler.NewStoreGRPCHandler(svc)
	svr := grpc.NewServer()
	pb.RegisterStoreGRPCHandlerServer(svr, tpt)
	if err := svr.Serve(host); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
