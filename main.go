package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"mohamedElsonny/simple-bank/api"
	db "mohamedElsonny/simple-bank/db/sqlc"
	"mohamedElsonny/simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".", "dev")
	if err != nil {
		log.Fatal("error on loading config file", err)
	}
	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("error on connecting database", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(config.Addr)
	if err != nil {
		log.Fatal("error on running server", err)
	}

}
