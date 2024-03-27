package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type Config struct {
	DB *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	app := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}