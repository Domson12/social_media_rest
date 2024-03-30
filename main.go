package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/Domson12/social_media_rest/api"
	db "github.com/Domson12/social_media_rest/db/sqlc"
	"github.com/Domson12/social_media_rest/util"
)

func main() {
	config, err1 := util.LoadConfig(".")
	if err1 != nil {
		log.Fatal("cannot load config: ", err1)
	}

	var err error
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
