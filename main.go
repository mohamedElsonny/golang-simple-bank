package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"mohamedElsonny/simple-bank/api"
	db "mohamedElsonny/simple-bank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:postgres@localhost:5555/simple_bank?sslmode=disable"
	addr     = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("error on connecting database", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(addr)
	if err != nil {
		log.Fatal("error on running server", err)
	}

}
