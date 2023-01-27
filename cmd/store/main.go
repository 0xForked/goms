package main

import (
	"database/sql"
	"fmt"
	"github.com/aasumitro/goms/internal/store"
	"net"
	"sync"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

const (
	DBDriver       = "sqlite3"
	DBSource       = "./db/store.db"
	ServiceNetwork = "tcp"
	ServiceAddress = "localhost:8001"
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
	if listener, err = net.Listen(ServiceNetwork, ServiceAddress); err != nil {
		panic(fmt.Sprintf(
			"LISTENER_ERROR: %s",
			err.Error()))
	}
	defer func() { _ = listener.Close() }()
	store.NewStoreService(dbPool, listener)
}

func getDBConn() {
	dbOnce.Do(func() {
		if dbPool, err = sql.Open(DBDriver, DBSource); err != nil {
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
