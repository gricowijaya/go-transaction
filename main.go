package main

import (
	"database/sql"
	"log"

	"github.com/gricowijaya/go-transaction/api"
	db "github.com/gricowijaya/go-transaction/db/sqlc"
	"github.com/gricowijaya/go-transaction/util"
	_ "github.com/lib/pq"
)


func main() {
  config, err := util.LoadConfig(".")
  if err != nil{ 
    log.Fatal("Cannot check Config")
  }
  // connect to the database
  conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot Connect to the Database because ", err)
	}

  store := db.NewStore(conn)
  server := api.NewServer(store)

  err = server.Start(config.ServerAddress)
  if err != nil {
    log.Fatal("Cannot start server -> ", err)
  }
}
