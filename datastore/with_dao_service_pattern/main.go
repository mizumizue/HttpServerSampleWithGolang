package main

import (
	"go-server-samples/datastore/with_dao_service_pattern/router"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := router.NewRouter()
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8888", router))
}
