
package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	} else {
		fmt.Println("ENV variables loaded")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	} else {
		fmt.Println("Connected to DB")
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal("Cannot create server", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
	