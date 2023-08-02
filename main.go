package main

import (
	"database/sql"
	"log"

	api "github.com/almacitunaberk/go_masterclass/api"
	db "github.com/almacitunaberk/go_masterclass/db/sqlc"
	"github.com/almacitunaberk/go_masterclass/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config file: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connec to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create the server: ", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start the server: ", err)
	}
}
