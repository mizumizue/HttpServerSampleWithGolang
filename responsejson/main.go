package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Name string		`json:"name"`
	Age int			`json:"age"`
	Sex string		`json:"sex"`
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	person := Person{"YamadaTarou", 24, "male"}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}
