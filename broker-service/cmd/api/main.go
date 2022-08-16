package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	//One root and responde json
	app := Config{}

	log.Printf("Broker Service has started at port: %s!\n", webPort)

	// Defining http servers
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// Starting the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
