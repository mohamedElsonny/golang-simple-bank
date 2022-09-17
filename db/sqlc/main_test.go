package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"mohamedElsonny/simple-bank/util"
)

var testQueries *Queries

var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../", "test")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	testDB, err = sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
