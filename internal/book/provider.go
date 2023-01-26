package store

import (
	"database/sql"
	handler "github.com/bakode/goms/internal/book/delivery/handler/grpc"
	sqlRepo "github.com/bakode/goms/internal/book/repository/sql"
	"github.com/bakode/goms/internal/book/service"
	"github.com/bakode/goms/pkg/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewBookService(db *sql.DB, host net.Listener) {
	repo := sqlRepo.NewBookSQLRepository(db)
	svc := service.NewBookService(repo)
	tpt := handler.NewBookGRPCHandler(svc)
	svr := grpc.NewServer()
	pb.RegisterBookGRPCHandlerServer(svr, tpt)
	if err := svr.Serve(host); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
