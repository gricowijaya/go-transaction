package db

import (
	"database/sql"
	_ "github.com/lib/pq"
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
  var err error
	testDB, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("Cannot Connect to the Database because ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
