package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/gricowijaya/go-transaction/util"
	"log"
	"os"
	"testing"
)

// the database driver and the source url from the docker in the makefile
const (
	driverName     = "postgres"
	dataSourceName = "postgresql://postgres:password@localhost:5000/user_golang?sslmode=disable"
)

// testQueries is a pointer to *Queries struct in the ./db/sqlc/db.go
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
  config, err := util.LoadConfig("../../.")
  if err !=nil {
    return
  }
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot Connect to the Database because ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
