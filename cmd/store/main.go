package main

import (
	"database/sql"
	"fmt"
	"github.com/bakode/goms/internal/store"
	"log"
	"net"
	"sync"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

const (
	DBDriver       = "sqlite3"
	DBSource       = "./db/store.db"
	ServiceNetwork = "tcp"
	ServiceAddress = ":8001"
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
		log.Fatalf("Could not listen on port: %v", err)
	}
	defer func() { _ = listener.Close() }()
	store.NewStoreService(dbPool, listener)
}

func getDBConn() {
	dbOnce.Do(func() {
		var err error

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
