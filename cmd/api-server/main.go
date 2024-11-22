package main

import (
	"log"
	"net/http"
	"newsAPI/cmd/api-server/internal/router"
)

func main() {

	r := router.New()
	//Building a HTTP server
	if err := http.ListenAndServe(":8080", r); err != nil {

		log.Fatal("Failed to start server ", err)

	}

}
