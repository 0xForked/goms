package main

import (
	"database/sql"
	"fmt"
	store "github.com/aasumitro/goms/internal/book"
	"net"
	"sync"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriver       = "sqlite3"
	dbSource       = "./db/book.db"
	serviceNetwork = "tcp"
	serviceAddress = "localhost:8002"
)

var (
	dbOnce   sync.Once
	dbPool   *sql.DB
	listener net.Listener
	err      error
)

func init() {
	getDBConn()
}

func main() {
	if listener, err = net.Listen(serviceNetwork, serviceAddress); err != nil {
		panic(fmt.Sprintf(
			"LISTENER_ERROR: %s",
			err.Error()))
	}
	defer func() { _ = listener.Close() }()
	store.NewBookService(dbPool, listener)
}

func getDBConn() {
	dbOnce.Do(func() {
		if dbPool, err = sql.Open(dbDriver, dbSource); err != nil {
			panic(fmt.Sprintf(
				"DATABASE_ERROR: %s",
				err.Error()))
		}

		if err = dbPool.Ping(); err != nil {
			panic(fmt.Sprintf(
				"DATABASE_ERROR: %s",
				err.Error()))
		}
	})
}
