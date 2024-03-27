package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080" //8080

type Config struct{}

func main() {
	app := Config{}

	log.Printf("starting  broker service on port %s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
