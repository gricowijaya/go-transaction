package main

import (
	"database/sql"
	"log"

	db "github.com/gricowijaya/go-transaction/db/sqlc"
	"github.com/gricowijaya/go-transaction/api"
	_ "github.com/lib/pq"
)

const (
	driverName     = "postgres"
	dataSourceName = "postgresql://postgres:password@localhost:5000/user_golang?sslmode=disable"
  serverAddress = "0.0.0.0:3940"
)

func main() {
  // connect to the database
  conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("Cannot Connect to the Database because ", err)
	}

  store := db.NewStore(conn)
  server := api.NewServer(store)

  err = server.Start(serverAddress)
  if err != nil {
    log.Fatal("Cannot start server -> ", err)
  }
}
