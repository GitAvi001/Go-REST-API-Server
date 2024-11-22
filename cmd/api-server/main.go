package main

import (
	"log"
	"net/http"
)

func main() {

	//Building a HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {

		log.Fatal("Failed to start server ", err)

	}

}
